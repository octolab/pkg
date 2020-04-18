package bytes

import (
	"strconv"
	"strings"
	"unicode"

	"go.octolab.org/errors"
)

const (
	// ErrInvalidByteQuantity is returned by the ParseSize
	// when the passed string is invalid.
	ErrInvalidByteQuantity errors.Message = "byte quantity must be a positive integer with a unit of measurement like M, MB, MiB, etc"

	_ = 1 << (10 * iota)
	kilobyte
	megabyte
	gigabyte
	terabyte
	petabyte
	exabyte
)

// A Size represents a size of digital information.
type Size uint64

// String returns a human-readable byte string of the form 10M,
// 12.5K, and so forth. The following units are available:
//  E: Exabyte
//  P: Petabyte
//  T: Terabyte
//  G: Gigabyte
//  M: Megabyte
//  K: Kilobyte
//  B: Byte
// The unit that results in the smallest number greater than
// or equal to 1 is always chosen.
func (size Size) String() string {
	var unit rune
	value := float64(size)

	switch {
	case size >= exabyte:
		unit = 'E'
		value = value / exabyte
	case size >= petabyte:
		unit = 'P'
		value = value / petabyte
	case size >= terabyte:
		unit = 'T'
		value = value / terabyte
	case size >= gigabyte:
		unit = 'G'
		value = value / gigabyte
	case size >= megabyte:
		unit = 'M'
		value = value / megabyte
	case size >= kilobyte:
		unit = 'K'
		value = value / kilobyte
	default:
		unit = 'B'
	}

	return strings.TrimSuffix(
		strconv.FormatFloat(value, 'f', 1, 64),
		".0",
	) + string(unit)
}

// ParseSize parses a string formatted by Size as bytes.
// Note binary-prefixed and SI prefixed units both mean a base-2 units
//  KB = K = KiB = 1024
//  MB = M = MiB = 1024 * K
//  GB = G = GiB = 1024 * M
//  TB = T = TiB = 1024 * G
//  PB = P = PiB = 1024 * T
//  EB = E = EiB = 1024 * P
func ParseSize(s string) (Size, error) {
	s = strings.ToUpper(strings.TrimSpace(s))

	i := strings.IndexFunc(s, unicode.IsLetter)
	if i == -1 {
		return 0, ErrInvalidByteQuantity
	}

	raw, multiple := s[:i], s[i:]
	bytes, err := strconv.ParseFloat(raw, 64)
	if err != nil || bytes < 0 {
		return 0, ErrInvalidByteQuantity
	}

	var size Size
	switch multiple {
	case "E", "EB", "EIB":
		size = Size(bytes * exabyte)
	case "P", "PB", "PIB":
		size = Size(bytes * petabyte)
	case "T", "TB", "TIB":
		size = Size(bytes * terabyte)
	case "G", "GB", "GIB":
		size = Size(bytes * gigabyte)
	case "M", "MB", "MIB":
		size = Size(bytes * megabyte)
	case "K", "KB", "KIB":
		size = Size(bytes * kilobyte)
	case "B":
		size = Size(bytes)
	default:
		return 0, ErrInvalidByteQuantity
	}
	return size, nil
}
