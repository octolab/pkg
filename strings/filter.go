package strings

func Exclude(from []string, what ...string) []string {
	// if len(from) > x && len(what) > y { map }
	filtered := from[:0]
	for _, row := range from {
		found := false
		for _, target := range what {
			if row == target {
				found = true
				break
			}
		}
		if !found {
			filtered = append(filtered, row)
		}
	}
	return filtered
}

// FirstNotEmpty returns a first non-empty string.
func FirstNotEmpty(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

// NotEmpty filters empty strings in-place.
func NotEmpty(strings []string) []string {
	filtered := strings[:0]
	for _, str := range strings {
		if str != "" {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

// Unique filters non-unique strings in-place.
func Unique(strings []string) []string {
	registry := map[string]struct{}{}
	filtered := strings[:0]
	for _, str := range strings {
		if _, present := registry[str]; present {
			continue
		}
		registry[str] = struct{}{}
		filtered = append(filtered, str)
	}
	return filtered
}
