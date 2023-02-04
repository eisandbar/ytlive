export interface Stream {
  // Id: The ID that YouTube uses to uniquely identify the video.
  id?: string;

  // ActualStartTime: The time that the broadcast actually started. This
  // value will not be available until the broadcast begins.
  actualStartTime?: string;

  // ConcurrentViewers: The number of viewers currently watching the
  // broadcast. The property and its value will be present if the
  // broadcast has current viewers and the broadcast owner has not hidden
  // the viewcount for the video. Note that YouTube stops tracking the
  // number of concurrent viewers for a broadcast when the broadcast ends.
  // So, this property would not identify the number of viewers watching
  // an archived video of a live broadcast that already ended.
  concurrentViewers?: number;

  // CategoryId: The YouTube video category associated with the video.
  categoryId?: string;

  // The game being streamed, if applicable
  category?: string;

  // ChannelId: The ID that YouTube uses to uniquely identify the channel
  // that the video was uploaded to.
  channelId?: string;

  // ChannelTitle: Channel title for the channel that the video belongs
  // to.
  channelTitle?: string;

  // DefaultAudioLanguage: The default_audio_language property specifies
  // the language spoken in the video's default audio track.
  defaultAudioLanguage?: string;

  // DefaultLanguage: The language of the videos's default snippet.
  defaultLanguage?: string;

  // Description: The video's description. @mutable youtube.videos.insert
  // youtube.videos.update
  description?: string;

  // LiveBroadcastContent: Indicates if the video is an upcoming/active
  // live broadcast. Or it's "none" if the video is not an upcoming/active
  // live broadcast.
  //
  // Possible values:
  //   "none"
  //   "upcoming" - The live broadcast is upcoming.
  //   "live" - The live broadcast is active.
  //   "completed" - The live broadcast has been completed.
  liveBroadcastContent?: string;

  // PublishedAt: The date and time when the video was uploaded.
  publishedAt?: string;
  // Tags: A list of keyword tags associated with the video. Tags may
  // contain spaces.
  tags?: string[];

  // Thumbnails: A map of thumbnail images associated with the video. For
  // each object in the map, the key is the name of the thumbnail image,
  // and the value is an object that contains other information about the
  // thumbnail.
  thumbnail?: string;

  // Title: The video's title. @mutable youtube.videos.insert
  // youtube.videos.update
  title?: string;
}
