package runtime_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/runtime"
)

func TestVersion_Compare(t *testing.T) {
	bump := func(version GoVersion) GoVersion {
		version.Minor++
		return version
	}

	version := Version()
	if unstable(version.Raw) {
		version.Major, version.Minor, version.Patch, version.Raw = 1, 12, 0, "go1.12"
	}

	tests := map[string]struct {
		target  GoVersion
		compare func(GoVersion, GoVersion) bool
	}{
		"before":     {GoVersion{Major: 2}, GoVersion.Before},
		"closely":    {bump(version), GoVersion.Before},
		"later":      {GoVersion{Major: 1}, GoVersion.Later},
		"much later": {GoVersion{}, GoVersion.Later},
		"equal":      {version, GoVersion.Equal},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.True(t, test.compare(version, test.target))
		})
	}
	t.Run("full comparison", func(t *testing.T) {
		base := GoVersion{Major: 1, Minor: 9}
		v1, v2, v3 := base, base, base
		v1.Patch, v2.Patch, v3.Patch = 1, 2, 3
		assert.True(t, v1.Before(v2) && v2.Before(v3))
		assert.True(t, v3.Later(v2) && v2.Later(v1))
		assert.True(t, v2.Later(v1) && v2.Before(v3))
		assert.True(t, !base.Later(base) && !base.Before(base))
	})
}

func unstable(version string) bool {
	return strings.HasPrefix(version, "devel")
}

var go112 = struct {
	version GoVersion
	release string
}{
	version: GoVersion{Major: 1, Minor: 12, Raw: "go1.12"},
	release: "Mon Feb 25 16:47:57 2019 -0500",
}
