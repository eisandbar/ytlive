package store

import (
	pq "github.com/lib/pq"
	"google.golang.org/api/youtube/v3"
)

func StreamData(video youtube.Video) Stream {
	stream := Stream{
		Id:                   video.Id,
		ActualStartTime:      video.LiveStreamingDetails.ActualStartTime,
		ConcurrentViewers:    video.LiveStreamingDetails.ConcurrentViewers,
		CategoryId:           video.Snippet.CategoryId,
		ChannelId:            video.Snippet.ChannelId,
		ChannelTitle:         video.Snippet.ChannelTitle,
		DefaultAudioLanguage: video.Snippet.DefaultAudioLanguage,
		DefaultLanguage:      video.Snippet.DefaultLanguage,
		Description:          video.Snippet.Description,
		LiveBroadcastContent: video.Snippet.LiveBroadcastContent,
		Tags:                 video.Snippet.Tags,
		Title:                video.Snippet.Title,
		Thumbnail:            video.Snippet.Thumbnails.Medium.Url,
	}
	return stream
}

type Stream struct {
	// Id: The ID that YouTube uses to uniquely identify the video.
	Id string `json:"id,omitempty" gorm:"primaryKey"`

	// ActualStartTime: The time that the broadcast actually started. This
	// value will not be available until the broadcast begins.
	ActualStartTime string `json:"actualStartTime,omitempty"`

	// ConcurrentViewers: The number of viewers currently watching the
	// broadcast. The property and its value will be present if the
	// broadcast has current viewers and the broadcast owner has not hidden
	// the viewcount for the video. Note that YouTube stops tracking the
	// number of concurrent viewers for a broadcast when the broadcast ends.
	// So, this property would not identify the number of viewers watching
	// an archived video of a live broadcast that already ended.
	ConcurrentViewers uint64 `json:"concurrentViewers,omitempty,string"`

	// CategoryId: The YouTube video category associated with the video.
	CategoryId string `json:"categoryId,omitempty"`

	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// that the video was uploaded to.
	ChannelId string `json:"channelId,omitempty"`

	// ChannelTitle: Channel title for the channel that the video belongs
	// to.
	ChannelTitle string `json:"channelTitle,omitempty"`

	// DefaultAudioLanguage: The default_audio_language property specifies
	// the language spoken in the video's default audio track.
	DefaultAudioLanguage string `json:"defaultAudioLanguage,omitempty"`

	// DefaultLanguage: The language of the videos's default snippet.
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// Description: The video's description. @mutable youtube.videos.insert
	// youtube.videos.update
	Description string `json:"description,omitempty"`

	// LiveBroadcastContent: Indicates if the video is an upcoming/active
	// live broadcast. Or it's "none" if the video is not an upcoming/active
	// live broadcast.
	//
	// Possible values:
	//   "none"
	//   "upcoming" - The live broadcast is upcoming.
	//   "live" - The live broadcast is active.
	//   "completed" - The live broadcast has been completed.
	LiveBroadcastContent string `json:"liveBroadcastContent,omitempty"`

	// PublishedAt: The date and time when the video was uploaded.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: A list of keyword tags associated with the video. Tags may
	// contain spaces.
	Tags pq.StringArray `json:"tags,omitempty" gorm:"type:text[]"`

	// Thumbnails: A map of thumbnail images associated with the video. For
	// each object in the map, the key is the name of the thumbnail image,
	// and the value is an object that contains other information about the
	// thumbnail.
	Thumbnail string `json:"thumbnail,omitempty"`

	// Title: The video's title. @mutable youtube.videos.insert
	// youtube.videos.update
	Title string `json:"title,omitempty"`

	// The category of the stream
	Category string `json:"category,omitempty"`
}
