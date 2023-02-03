package adder

import (
	"log"
	"time"

	"github.com/eisandbar/ytlive/app/internal"
	"github.com/eisandbar/ytlive/app/store"
	"github.com/eisandbar/ytlive/app/youtubeapi"
	"google.golang.org/api/youtube/v3"
)

func NewAdder(ds store.Store, service *youtube.Service, gaming bool, interval, timeMissed time.Duration) {
	adder := Adder{}
	adder.store = ds
	adder.api = youtubeapi.NewYoutubeAPI(service)

	// Catchup for the time the server was offline
	catchup(timeMissed, gaming, adder.AddStreams)

	addStreams := func() {
		so := &youtubeapi.SearchOptions{}
		so.WithPublishedAfter(string(time.Now().Add(-interval - time.Minute).Format(time.RFC3339)))
		so.WithPublishedBefore(time.Now().Format(time.RFC3339))
		if gaming {
			so.Gaming()
		}
		adder.AddStreams(so)
	}

	go internal.Ticker(interval, addStreams)
}

type Adder struct {
	store store.Store
	api   youtubeapi.API
}

func (adder *Adder) AddStreams(options *youtubeapi.SearchOptions) {
	log.Println("Adding streams")
	// Get new list of streams
	searchList := adder.api.GetSearchList(options)

	log.Printf("Found %d of %d streams", searchList.PageInfo.ResultsPerPage, searchList.PageInfo.TotalResults)
	log.Println(searchList.PrevPageToken, searchList.NextPageToken)
	videoList := adder.api.GetVideosList(youtubeapi.GetVideoId(searchList))
	for _, video := range videoList.Items {
		stream := store.StreamData(*video)
		adder.store.Add(stream)
	}

	if searchList.NextPageToken != "" {
		options.WithToken(searchList.NextPageToken)
		adder.AddStreams(options)
	}
}
