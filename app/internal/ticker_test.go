package internal_test

import (
	"testing"
	"time"

	"github.com/eisandbar/ytlive/app/internal"
	"github.com/stretchr/testify/assert"
)

func TestTicker(t *testing.T) {
	value := 0
	addOne := func() {
		value++
	}
	go internal.Ticker(time.Millisecond*10, addOne)
	time.Sleep(time.Millisecond * 10)
	assert.Equal(t, 1, value)
	time.Sleep(time.Millisecond * 11)
	assert.Equal(t, 2, value)
	time.Sleep(time.Millisecond * 11)
	assert.Equal(t, 3, value)
}

func TestTickerCommunication(t *testing.T) {
	intChan := make(chan int, 5)
	value := 0
	addOne := func() {
		value += <-intChan
	}

	go internal.Ticker(time.Millisecond*10, addOne)

	values := []int{2, 4, 7}
	sums := []int{2, 6, 13}

	for i, v := range values {
		intChan <- v
		time.Sleep(time.Millisecond * 11)
		assert.Equal(t, sums[i], value)
	}
}
