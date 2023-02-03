package store

type Category struct {
	Category          string `json:"category,omitempty" gorm:"primaryKey"`
	ConcurrentViewers uint64 `json:"concurrentViewers,omitempty,string"`
	Url               string `json:"url,omitempty"`
}
