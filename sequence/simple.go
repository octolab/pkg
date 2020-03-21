package sequence

// Simple returns an empty slice with the specified size.
//
//  for range Simple(5) {
//  	// do something five times
//  }
//
// Read the https://dave.cheney.net/2014/03/25/the-empty-struct
// and the https://github.com/bradfitz/iter.
func Simple(size int) []struct{} {
	return make([]struct{}, size)
}
