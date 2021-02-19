package testing

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.octolab.org/safe"
)

func Content(t *testing.T, urn string) []byte {
	t.Helper()

	// tmp solution to skip complexity
	urn = strings.TrimPrefix(urn, "file:")

	f, err := os.Open(urn)
	require.NoError(t, err)
	defer safe.Close(f, NoError(t))

	d, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	return d
}
