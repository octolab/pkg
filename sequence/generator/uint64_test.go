package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/sequence/generator"
)

func TestUint64(t *testing.T) {
	generator := new(Uint64).At(10)

	assert.Equal(t, uint64(10), generator.Current())
	assert.Equal(t, uint64(11), generator.Next())
	assert.Equal(t, uint64(21), generator.Jump(10))
	assert.Equal(t, uint64(21), generator.Reset())
	assert.Equal(t, uint64(0), generator.Current())
}
