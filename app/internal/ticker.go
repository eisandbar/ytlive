package internal

import "time"

func Ticker(interval time.Duration, callback func()) {
	ticker := time.NewTicker(interval)
	for {
		_, next := <-ticker.C
		if next {
			callback()
		} else {
			continue
		}
	}
}
