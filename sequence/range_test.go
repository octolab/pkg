package sequence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "go.octolab.org/sequence"
)

func TestRange(t *testing.T) {
	tests := map[string]struct {
		size, batch int
		expected    []int
	}{
		"empty":    {0, 7, []int{}},
		"less":     {5, 7, []int{5}},
		"equal":    {7, 7, []int{7}},
		"greater":  {10, 7, []int{7, 10}},
		"multiply": {21, 7, []int{7, 14, 21}},
		"common":   {17, 7, []int{7, 14, 17}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, Range(test.size, test.batch))
		})
	}
}
