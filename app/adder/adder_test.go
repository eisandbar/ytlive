package adder_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	a := time.Minute * 100
	a = time.Duration(float64(a) * 1.2)
	assert.Equal(t, time.Minute*120, a)
}

// Even though they are declared i the same line, one is greater than the other
func TestAfter(t *testing.T) {
	a, b := time.Now(), time.Now()
	assert.Equal(t, true, b.After(a))
}
