package store_test

import (
	"testing"

	"github.com/eisandbar/ytlive/app/store"
	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {

	pgStore := store.NewPGStore()
	res := pgStore.Categories(true)
	// spew.Dump(res[0:10])
	assert.Greater(t, int(res[0].ConcurrentViewers), 0)
}

func TestStuff(t *testing.T) {
	test := Test{}
	test.applyA("A")
	test.applyB("B")
	assert.Equal(t, "A", test.A)
	assert.Equal(t, "B", test.B)
}

type Test struct {
	A string
	B string
}

func (t *Test) applyA(text string) {
	t.A = text
}
func (t *Test) applyB(text string) {
	t.B = text
}
