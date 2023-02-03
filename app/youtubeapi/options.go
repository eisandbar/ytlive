package youtubeapi

type SearchOptions struct {
	token           string
	publishedAfter  string
	publishedBefore string
	gaming          bool
}

func (so *SearchOptions) WithPublishedAfter(time string) {
	so.publishedAfter = time
}

func (so *SearchOptions) WithPublishedBefore(time string) {
	so.publishedBefore = time
}

func (so *SearchOptions) WithToken(token string) {
	so.token = token
}

func (so *SearchOptions) Gaming() {
	so.gaming = true
}
