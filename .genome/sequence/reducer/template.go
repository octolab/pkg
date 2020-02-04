// +build genome

package reducer

import (
	"errors"
	"sort"
)

// Interface defines behavior of a generic reducer above a numeric sequence.
type Interface interface {
	// Average returns an average value of the sequence.
	Average() float64
	// Length returns the sequence length.
	Length() int
	// Maximum returns a maximum value of the sequence.
	Maximum() T
	// Median returns a median value of the sequence.
	Median() float64
	// Minimum returns a minimum value of the sequence.
	Minimum() T
	// Sum returns a sum of the sequence.
	Sum() T
}

type T int

// ReduceT wraps sequence of integers to perform aggregate operations above it.
func ReduceT(sequence ...T) (Interface, error) {
	if len(sequence) == 0 {
		return nil, errors.New("convert me")
	}
	return reducer(sequence), nil
}

type reducer []T

// Average returns an average value of the sequence.
func (sequence reducer) Average() float64 {
	return float64(sequence.Sum()) / float64(len(sequence))
}

// Length returns the sequence length.
func (sequence reducer) Length() int {
	return len(sequence)
}

// Maximum returns a maximum value of the sequence.
func (sequence reducer) Maximum() T {
	max := sequence[0]
	for _, num := range sequence {
		if num > max {
			max = num
		}
	}
	return max
}

// Median returns a median value of the sequence.
func (sequence reducer) Median() float64 {
	size := len(sequence)
	sorted := append(make([]T, 0, size), sequence...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	if size%2 == 0 {
		return (float64(sorted[size/2-1]) + float64(sorted[size/2])) / 2
	}
	return float64(sorted[size/2])
}

// Minimum returns a minimum value of the sequence.
func (sequence reducer) Minimum() T {
	min := sequence[0]
	for _, num := range sequence {
		if num < min {
			min = num
		}
	}
	return min
}

// Sum returns a sum of the sequence.
func (sequence reducer) Sum() T {
	var sum T
	for _, num := range sequence {
		sum += num
	}
	return sum
}
