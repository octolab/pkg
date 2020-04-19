package sequence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/sequence"
)

func TestIntReducer_Average(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected float64
	}{
		"empty case":  {[]int{}, 0},
		"nil case":    {nil, 0},
		"normal case": {[]int{1, 2, 3}, 2},
		"fractional":  {[]int{1, 2, 3, 4, 5, 6}, 3.5},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Average())
		})
	}
}

func TestIntReducer_Length(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"normal case": {[]int{1, 2, 3}, 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Length())
		})
	}
}

func TestIntReducer_Maximum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"empty case": {[]int{}, 0},
		"nil case":   {nil, 0},
		"sorted":     {[]int{1, 2, 3}, 3},
		"unsorted":   {[]int{3, 2, 1}, 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Maximum())
		})
	}
}

func TestIntReducer_Median(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected float64
	}{
		"empty case":     {[]int{}, 0},
		"nil case":       {nil, 0},
		"even, sorted":   {[]int{1, 2, 3, 4}, 2.5},
		"even, unsorted": {[]int{2, 1, 4, 3}, 2.5},
		"odd, sorted":    {[]int{1, 2, 3, 4, 5}, 3},
		"odd, unsorted":  {[]int{3, 1, 2, 4, 5}, 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Median())
		})
	}
}

func TestIntReducer_Minimum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"empty case": {[]int{}, 0},
		"nil case":   {nil, 0},
		"sorted":     {[]int{1, 2, 3}, 1},
		"unsorted":   {[]int{3, 2, 1}, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Minimum())
		})
	}
}

func TestIntReducer_Sum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"empty case":      {[]int{}, 0},
		"nil case":        {nil, 0},
		"positive sum":    {[]int{1, 2, 3}, 6},
		"negative sum":    {[]int{-1, -2, -3}, -6},
		"mixed, zero sum": {[]int{-1, -2, 3}, 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, ReduceInts(test.sequence...).Sum())
		})
	}
}
