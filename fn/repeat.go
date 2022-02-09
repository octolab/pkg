package fn

// Repeat repeats the action the required number of times.
//
//	func FillByValue(slice []int, value int) error {
//		return fn.Repeat(
//			fn.HasNoError(func () { slice = append(slice, value) }),
//			cap(slice) - len(slice),
//		)
//	}
func Repeat(action func() error, times int) error {
	for i := 0; i < times; i++ {
		if err := action(); err != nil {
			return err
		}
	}
	return nil
}
