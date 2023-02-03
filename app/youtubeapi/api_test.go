package youtubeapi_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/eisandbar/ytlive/app/youtubeapi"
	"github.com/stretchr/testify/assert"
)

// This test shows that if a stream was delete or taken down, using the youtube api
// will return nothing for that stream. So of these 4 streams, 3 were taken down
// so only 1 result is returned
func TestMissingItems(t *testing.T) {
	ids := []string{"LcisZwjMrUg", "USvpG0dusxg", "VFqPKPC8Npc", "PQqRcLM1Deo"}

	service := youtubeapi.NewService()
	api := youtubeapi.NewYoutubeAPI(service)
	res := api.GetVideosList(ids)
	spew.Dump(res)
	assert.Equal(t, 1, len(res.Items))
}
