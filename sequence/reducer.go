package sequence

// Reducer defines behavior of a generic reducer above a numeric sequence.
type Reducer interface {
	// Average returns an average value of the sequence.
	Average() float64
	// Length returns the sequence length.
	Length() int
	// Maximum returns a maximum value of the sequence.
	Maximum() int
	// Median returns a median value of the sequence.
	Median() float64
	// Minimum returns a minimum value of the sequence.
	Minimum() int
	// Sum returns a sum of the sequence.
	Sum() int
}
