package strings

func PresentAny(where []string, what ...string) bool {
	// if len(where) > x && len(what) > y { map }
	for _, row := range where {
		for _, target := range what {
			if row == target {
				return true
			}
		}
	}
	return false
}
