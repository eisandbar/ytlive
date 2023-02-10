package adder

import (
	"time"

	"github.com/eisandbar/ytlive/app/youtubeapi"
)

// Exponentially increases interval to catchup to missed streams
func catchup(timeMissed time.Duration, gaming bool, callback func(so *youtubeapi.SearchOptions)) {
	current, now := time.Now(), time.Now()
	interval := time.Minute * 20

	for current.After(now.Add(-timeMissed)) {
		so := &youtubeapi.SearchOptions{}
		if gaming {
			so.Gaming()
		}
		so.WithPublishedBefore(current.Format(time.RFC3339))
		so.WithPublishedAfter(current.Add(-interval).Format(time.RFC3339))

		go callback(so)

		current = current.Add(-interval)
		interval = time.Duration(float64(interval) * 1.2)
	}
}
