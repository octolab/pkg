package sequence

import (
	"sort"

	"go.octolab.org/errors"
)

const InvalidSequence errors.Message = "invalid sequence"

// ReduceInts wraps sequence of integers to perform aggregate operations above it.
//
//  ticket, err := semaphore.Acquire(ctx, sequence.ReduceInts(places...).Sum())
//
//  if err != nil {
//  	...
//  }
//  defer semaphore.Release(ticket)
//
func ReduceInts(sequence ...int) (Reducer, error) {
	if len(sequence) == 0 {
		return nil, InvalidSequence
	}
	return intReducer(sequence), nil
}

type intReducer []int

// Average returns an average value of the sequence.
func (sequence intReducer) Average() float64 {
	return float64(sequence.Sum()) / float64(len(sequence))
}

// Length returns the sequence length.
func (sequence intReducer) Length() int {
	return len(sequence)
}

// Maximum returns a maximum value of the sequence.
func (sequence intReducer) Maximum() int {
	max := sequence[0]
	for _, num := range sequence {
		if num > max {
			max = num
		}
	}
	return max
}

// Median returns a median value of the sequence.
func (sequence intReducer) Median() float64 {
	size := len(sequence)
	sorted := append(make([]int, 0, size), sequence...)
	sort.Ints(sorted)
	if size%2 == 0 {
		return (float64(sorted[size/2-1]) + float64(sorted[size/2])) / 2
	}
	return float64(sorted[size/2])
}

// Minimum returns a minimum value of the sequence.
func (sequence intReducer) Minimum() int {
	min := sequence[0]
	for _, num := range sequence {
		if num < min {
			min = num
		}
	}
	return min
}

// Sum returns a sum of the sequence.
func (sequence intReducer) Sum() int {
	sum := 0
	for _, num := range sequence {
		sum += num
	}
	return sum
}
