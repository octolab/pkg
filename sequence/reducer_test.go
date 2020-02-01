package sequence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "go.octolab.org/sequence"
)

func TestIntReducer(t *testing.T) {
	tests := map[string][]int{
		"nil, invalid":   nil,
		"empty, invalid": {},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test...)
			assert.Nil(t, reducer)
			assert.Error(t, err)
			assert.True(t, err == InvalidSequence)
		})
	}
}

func TestIntReducer_Average(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected float64
	}{
		"normal case": {[]int{1, 2, 3}, 2},
		"fractional":  {[]int{1, 2, 3, 4, 5, 6}, 3.5},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Average())
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
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Length())
		})
	}
}

func TestIntReducer_Maximum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"sorted":   {[]int{1, 2, 3}, 3},
		"unsorted": {[]int{3, 2, 1}, 3},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Maximum())
		})
	}
}

func TestIntReducer_Median(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected float64
	}{
		"even, sorted":   {[]int{1, 2, 3, 4}, 2.5},
		"even, unsorted": {[]int{2, 1, 4, 3}, 2.5},
		"odd, sorted":    {[]int{1, 2, 3, 4, 5}, 3},
		"odd, unsorted":  {[]int{3, 1, 2, 4, 5}, 3},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Median())
		})
	}
}

func TestIntReducer_Minimum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"sorted":   {[]int{1, 2, 3}, 1},
		"unsorted": {[]int{3, 2, 1}, 1},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Minimum())
		})
	}
}

func TestIntReducer_Sum(t *testing.T) {
	tests := map[string]struct {
		sequence []int
		expected int
	}{
		"positive sum":    {[]int{1, 2, 3}, 6},
		"negative sum":    {[]int{-1, -2, -3}, -6},
		"mixed, zero sum": {[]int{-1, -2, 3}, 0},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reducer, err := ReduceInts(test.sequence...)
			assert.Nil(t, err)
			require.NotNil(t, reducer)
			assert.Equal(t, test.expected, reducer.Sum())
		})
	}
}
