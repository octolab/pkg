package numd

// Decline returns the word corresponding to the number.
func Decline(number int, words ...string) string {
	length := len(words)
	if length < 2 {
		return ""
	}

	if number < 0 {
		number = -number
	}

	switch {
	case number != 1 && length == 2:
		return words[1]
	case isGenitivePlural(number):
		return words[2]
	case isGenitiveSingular(number):
		return words[1]
	default:
		return words[0]
	}
}

// Genitive plural test.
func isGenitivePlural(num int) bool {
	nn := num % 10
	return (num > 10 && ((num%100)-((num%100)%10))/10 == 1) || (nn == 0 || nn >= 5)
}

// Genitive singular test.
func isGenitiveSingular(num int) bool {
	return num%10 >= 2
}
