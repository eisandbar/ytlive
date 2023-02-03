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
