package youtubeapi

import (
	"log"

	"google.golang.org/api/youtube/v3"
)

type API interface {
	GetSearchList(options *SearchOptions) youtube.SearchListResponse
	GetVideosList(id []string) youtube.VideoListResponse
}

func NewYoutubeAPI(service *youtube.Service) *YoutubeAPI {
	youtubeAPI := YoutubeAPI{}
	youtubeAPI.service = service
	return &youtubeAPI
}

type YoutubeAPI struct {
	service *youtube.Service
}

func (yt *YoutubeAPI) GetSearchList(options *SearchOptions) youtube.SearchListResponse {
	listCall := BaseListCall(yt.service)
	listCall.PageToken(options.token)
	listCall.PublishedAfter(options.publishedAfter)
	listCall.PublishedBefore(options.publishedBefore)
	if options.gaming {
		listCall.VideoCategoryId("20")
		listCall.Q("gaming")
	}

	// Executing search
	listResponse, err := listCall.Do()
	if err != nil {
		log.Fatalf("Unable to get search list response: %v", err)
	}

	return *listResponse
}

func (yt *YoutubeAPI) GetVideosList(id []string) youtube.VideoListResponse {
	if len(id) == 0 {
		return youtube.VideoListResponse{}
	}

	// Creating new search service
	videoService := youtube.NewVideosService(yt.service)

	// Setting search params
	listCall := videoService.List([]string{"snippet", "liveStreamingDetails", "topicDetails"})
	listCall.Id(id...)

	// Executing search
	listResponse, err := listCall.Do()
	if err != nil {
		log.Fatalf("Unable to get video list response: %v", err)
	}
	return *listResponse
}

func BaseListCall(service *youtube.Service) *youtube.SearchListCall {

	// Creating new search service
	searchService := youtube.NewSearchService(service)

	listCall := searchService.List([]string{"snippet"})
	applySearchOptions(listCall)

	return listCall
}

func applySearchOptions(listCall *youtube.SearchListCall) {
	listCall.EventType("live") // only active broadcasts
	listCall.MaxResults(50)    // 50 max
	listCall.Type("video")     // has to be set for eventType = live
	listCall.Order("date")     // reverse chronological order
}

func GetVideoId(listResponse youtube.SearchListResponse) []string {
	videos := make([]string, len(listResponse.Items))
	for i := range videos {
		videos[i] = listResponse.Items[i].Id.VideoId
	}

	return videos
}
