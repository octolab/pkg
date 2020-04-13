package fn_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/fn"
)

func TestRepeat(t *testing.T) {
	tests := map[string]int{
		"zero times":     0,
		"constant times": 5,
		"random times":   rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000),
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var counter int
			Repeat(func() { counter++ }, test)
			assert.Equal(t, test, counter)
		})
	}
}
