package testing

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/env"
)

func Test_setEnvs(t *testing.T) {
	os.Clearenv()

	release, err := setEnvs(
		func(string) (string, bool) { return "", false },
		func(string, string) error { return errors.New("unhealthy") },
		func(string) error { return nil },
		func(error) {},
		env.Must(env.GoArch, "amd64"),
		env.Must(env.GoDebug, "allocfreetrace=1,clobberfree=1"),
		env.Must(env.GoGC, "off"),
		env.Must(env.GoMaxProcs, "1"),
		env.Must(env.GoOS, "darwin"),
		env.Must(env.GoPath, "/home/dir"),
		env.Must(env.GoRace, "1"),
		env.Must(env.GoRoot, "/usr/bin"),
		env.Must(env.GoTraceback, "system"),
	)
	assert.Error(t, err)
	assert.Nil(t, release)

	os.Clearenv()
}
