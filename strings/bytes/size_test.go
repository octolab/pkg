package bytes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/strings/bytes"
)

func TestParseSize(t *testing.T) {
	type expected struct {
		size Size
		err  error
	}

	tests := map[string]struct {
		size     string
		expected expected
	}{
		"Bytes":           {"10B", expected{10, nil}},
		"Kilobyte":        {"10K", expected{10 * Kilobyte, nil}},
		"Kilobyte, alt 1": {"10KB", expected{10 * Kilobyte, nil}},
		"Kilobyte, alt 2": {"10KiB", expected{10 * Kilobyte, nil}},
		"Megabyte":        {"10M", expected{10 * Megabyte, nil}},
		"Megabyte, alt 1": {"10MB", expected{10 * Megabyte, nil}},
		"Megabyte, alt 2": {"10MiB", expected{10 * Megabyte, nil}},
		"Gigabyte":        {"10G", expected{10 * Gigabyte, nil}},
		"Gigabyte, alt 1": {"10GB", expected{10 * Gigabyte, nil}},
		"Gigabyte, alt 2": {"10GiB", expected{10 * Gigabyte, nil}},
		"Terabyte":        {"10T", expected{10 * Terabyte, nil}},
		"Terabyte, alt 1": {"10TB", expected{10 * Terabyte, nil}},
		"Terabyte, alt 2": {"10TiB", expected{10 * Terabyte, nil}},
		"Petabyte":        {"10P", expected{10 * Petabyte, nil}},
		"Petabyte, alt 1": {"10PB", expected{10 * Petabyte, nil}},
		"Petabyte, alt 2": {"10PiB", expected{10 * Petabyte, nil}},
		"Exabyte":         {"10E", expected{10 * Exabyte, nil}},
		"Exabyte, alt 1":  {"10EB", expected{10 * Exabyte, nil}},
		"Exabyte, alt 2":  {"10EiB", expected{10 * Exabyte, nil}},
		"invalid value":   {",0B", expected{0, ErrInvalidByteQuantity}},
		"invalid unit":    {"10AiB", expected{0, ErrInvalidByteQuantity}},
		"without unit":    {"10", expected{0, ErrInvalidByteQuantity}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			size, err := ParseSize(test.size)
			assert.Equal(t, test.expected.size, size)
			assert.Equal(t, test.expected.err, err)
		})
	}
}

func TestSize(t *testing.T) {
	tests := map[string]struct {
		size     Size
		expected string
	}{
		"Bytes":           {Byte, "1B"},
		"Kilobyte":        {Kilobyte, "1K"},
		"Kilobyte, float": {1.5 * Kilobyte, "1.5K"},
		"Megabyte":        {Megabyte, "1M"},
		"Megabyte, float": {1.5 * Megabyte, "1.5M"},
		"Gigabyte":        {Gigabyte, "1G"},
		"Gigabyte, float": {1.5 * Gigabyte, "1.5G"},
		"Terabyte":        {Terabyte, "1T"},
		"Terabyte, float": {1.5 * Terabyte, "1.5T"},
		"Petabyte":        {Petabyte, "1P"},
		"Petabyte, float": {1.5 * Petabyte, "1.5P"},
		"Exabyte":         {Exabyte, "1E"},
		"Exabyte, float":  {1.5 * Exabyte, "1.5E"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.size.String())
		})
	}
}
