package internal_test

import (
	"os"
	"testing"

	"github.com/eisandbar/ytlive/app/internal"
	"github.com/stretchr/testify/assert"
)

func TestLoadToken(t *testing.T) {

	f, err := os.CreateTemp("", "sample")
	check(err)

	_, err = f.Write([]byte("Test"))
	check(err)

	token, err := internal.LoadToken(f.Name())
	check(err)

	assert.Equal(t, "Test", token)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
