package generator_test

import (
	"bytes"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/sequence/generator"
)

func TestUUID(t *testing.T) {
	generator, checks := new(UUID), 10<<5

	collision := make(map[string]struct{}, checks)
	for range make([]struct{}, checks) {
		id := generator.Next().String()
		require.NotEmpty(t, id)
		require.NotContains(t, collision, id)
		collision[id] = struct{}{}
	}
	require.Len(t, collision, checks)

	t.Run("bad case", func(t *testing.T) {
		uuid.SetRand(bytes.NewReader([]byte{0, 0}))
		id := generator.Next()
		require.Empty(t, id.String())
		require.False(t, id.Valid())
		uuid.SetRand(nil)
	})
}
