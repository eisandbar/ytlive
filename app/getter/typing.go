package getter

// {"contents":[{"richMetadataRenderer": // Start string
// },{"richMetadataRenderer": // End string
type RichMetadataRenderer struct {
	Thumbnail Thumbnail  `json:"thumbnail,omitempty"`
	Title     SimpleText `json:"title,omitempty"`
}

type SimpleText struct {
	SimpleText string `json:"simpleText,omitempty"`
}

type Thumbnail struct {
	Thumbnails []Thumbnails `json:"thumbnails,omitempty"`
}

type Thumbnails struct {
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
