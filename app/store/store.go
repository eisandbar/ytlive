package store

type Store interface {
	FindOne(Id string) Stream
	Add(video Stream)
	Delete(Id string)
	Update(video Stream)
	List(...ListOption) []Stream
	Categories(bool) []Category
	Len() int
	SaveCategory(category Category)
}
