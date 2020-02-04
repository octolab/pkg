package sequence

// Simple returns an empty slice with the specified size.
//
//  for range Simple(5) {
//  	// do something five times
//  }
//
func Simple(size int) []struct{} {
	return make([]struct{}, size)
}
