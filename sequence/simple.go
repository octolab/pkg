package sequence

// Simple returns an empty slice with the specified size.
//
//  for range Simple(5) {
//  	// do something five times
//  }
//
func Simple(size int) []int {
	return make([]int, size)
}
