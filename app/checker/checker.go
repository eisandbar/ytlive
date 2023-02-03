// Checker should query the data store for all streams and check their status every X seconds
// It should update the viewer count and delete streams that are no longer up (possibly with a buffer period)

package checker

import (
	"log"
	"time"

	"github.com/eisandbar/ytlive/app/internal"
	"github.com/eisandbar/ytlive/app/store"
	"github.com/eisandbar/ytlive/app/youtubeapi"
	"google.golang.org/api/youtube/v3"
)

const poolSize = 10

// NewChecker creates a checker that checks the streams every 'interval' seconds
func NewChecker(ds store.Store, service *youtube.Service, interval time.Duration) {
	checker := Checker{}
	checker.store = ds
	checker.api = youtubeapi.NewYoutubeAPI(service)
	// go checker.Check() // run once immediately because ticker doesn't

	go internal.Ticker(interval, checker.Check)
}

type Checker struct {
	store store.Store // data store
	api   youtubeapi.API
}

// Check queries the data store for a list of all streams and checks them
func (checker *Checker) Check() {
	log.Println("Checking streams")

	// get a list of all streams in the db
	count := checker.store.Len()

	// If the use of findOne puts too much strain on the db we can use a map here instead
	streams := checker.store.List(store.WithMaxResults(count))
	streamsMap := make(map[string]store.Stream, len(streams))
	for _, stream := range streams {
		streamsMap[stream.Id] = stream
	}

	// Create a buffer so only poolSize goroutines run at once
	pool := make(chan bool, poolSize)

	for i := 0; i < len(streams); i += 50 {

		ids := make([]string, 0, 50)
		for j := 0; j < 50 && i+j < len(streams); j++ {
			ids = append(ids, streams[i+j].Id)
		}

		check := func(ids []string, streamsMap map[string]store.Stream) {
			// Write to pool. Will be blocked if poolSize goroutines already running
			pool <- true

			// We need to check if we got a result for each stream
			didGetResult := make(map[string]bool)

			listResponse := checker.api.GetVideosList(ids)

			for _, video := range listResponse.Items {
				stream := store.StreamData(*video)

				// The stream we got from yt api has all fields updated other than the category
				stream.Category = streamsMap[stream.Id].Category

				didGetResult[stream.Id] = true

				if checkLive(stream) {
					checker.store.Update(stream)
				} else {
					// We want to keep a list of the streams for now
					// checker.store.Delete(stream.Id)
					checker.store.Update(stream)
				}
			}

			// streams we didn't get a result for were taken down, we can remove them from the db
			// or set live to none
			for _, id := range ids {
				if _, found := didGetResult[id]; !found {
					stream := streamsMap[id]
					stream.LiveBroadcastContent = "none"
					checker.store.Update(stream)
				}
			}

			// Free up space in the pool once finished
			<-pool
		}
		go check(ids, streamsMap)
	}
}

func checkLive(stream store.Stream) bool {
	return stream.LiveBroadcastContent == "live"
}
