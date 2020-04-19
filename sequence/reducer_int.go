package sequence

import "sort"

// ReduceInts wraps sequence of integers to perform aggregate operations above it.
//
//  ticket, err := semaphore.Acquire(ctx, sequence.ReduceInts(places...).Sum())
//
//  if err != nil {
//  	...
//  }
//  defer semaphore.Release(ticket)
//
func ReduceInts(sequence ...int) Reducer {
	return intReducer(sequence)
}

type intReducer []int

// Average returns an average value of the sequence.
func (sequence intReducer) Average() float64 {
	if len(sequence) == 0 {
		return 0
	}
	return float64(sequence.Sum()) / float64(len(sequence))
}

// Length returns the sequence length.
func (sequence intReducer) Length() int {
	return len(sequence)
}

// Maximum returns a maximum value of the sequence.
func (sequence intReducer) Maximum() int {
	if len(sequence) == 0 {
		return 0
	}
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
	if len(sequence) == 0 {
		return 0
	}
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
	if len(sequence) == 0 {
		return 0
	}
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
