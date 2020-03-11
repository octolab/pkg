package timer

import "time"

// Stopwatch calculates the fn execution time.
//
//  var result interface{}
//
//  duration := fn.Stopwatch(func() { result = do.some("heavy") })
//
func Stopwatch(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
