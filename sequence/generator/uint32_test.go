package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/sequence/generator"
)

func TestUint32(t *testing.T) {
	generator := new(Uint32).At(10)

	assert.Equal(t, uint32(10), generator.Current())
	assert.Equal(t, uint32(11), generator.Next())
	assert.Equal(t, uint32(21), generator.Jump(10))
	assert.Equal(t, uint32(21), generator.Reset())
	assert.Equal(t, uint32(0), generator.Current())
}
