package sequence_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	. "go.octolab.org/sequence"
)

func TestSimple(t *testing.T) {
	tests := map[string]int{
		"constant": 5,
		"random":   rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000),
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Len(t, Simple(test), test)
		})
	}
}
