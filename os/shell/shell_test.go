package shell_test

import (
	"errors"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/os/shell"
)

func TestClassify(t *testing.T) {
	type expected struct {
		sh  Shell
		err error
	}

	type test struct {
		name string
		bin  string
		ctx  Context
		skip bool
		expected
	}

	tests := []test{
		{
			"sh",
			"/bin/sh",
			All,
			false,
			expected{Sh, nil},
		},
		{
			"bash",
			"/bin/bash",
			All,
			false,
			expected{Bash, nil},
		},
		{
			"zsh",
			"/usr/local/bin/zsh",
			All,
			false,
			expected{Zsh, nil},
		},
		{
			"PowerShell",
			"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
			All,
			runtime.GOOS != "windows",
			expected{PowerShell, nil},
		},
		{
			"PowerShell hack",
			"/usr/local/bin/powershell.exe",
			All,
			false,
			expected{PowerShell, nil},
		},
		{
			"fish",
			"/usr/local/bin/fish",
			All,
			false,
			expected{err: errors.New(`shell: cannot classify shell by "/usr/local/bin/fish"`)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.skip {
				t.SkipNow()
			}

			sh, err := Classify(test.bin, test.ctx)

			if test.expected.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, test.expected.sh.String(), sh.String())
				return
			}
			assert.Error(t, err)
			assert.EqualError(t, err, test.expected.err.Error())
			assert.Empty(t, sh.String())
		})
	}

	assert.Panics(t, func() { _, _ = Classify("", All) })
}
