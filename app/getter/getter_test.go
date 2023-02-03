package getter_test

import (
	"testing"

	"github.com/eisandbar/ytlive/app/getter"
	"github.com/eisandbar/ytlive/app/store"
	"github.com/stretchr/testify/assert"
)

func TestGetCategory(t *testing.T) {
	id := ""
	category := "Minecraft"
	if id == "" {
		return
	}

	ts := TestStore{}
	myGetter := getter.Getter{}
	myGetter.Store = &ts

	stream := store.Stream{Id: id}
	myGetter.GetCategory(stream)
	assert.Equal(t, category, ts.category)
}

type TestStore struct {
	category string
}

func (ts *TestStore) FindOne(Id string) store.Stream {
	return store.Stream{}
}

func (ts *TestStore) Add(video store.Stream) {
}

func (ts *TestStore) Delete(Id string) {
}

func (ts *TestStore) Update(video store.Stream) {
	ts.category = video.Category
}

func (ts *TestStore) List(...store.ListOption) []store.Stream {
	return []store.Stream{}
}

func (ts *TestStore) Len() int {
	return 0
}

func (ts *TestStore) Categories(bool) []store.Category {
	return nil
}

func (ts *TestStore) SaveCategory(category store.Category) {
}
