package numd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/strings/numd"
)

func TestDecline(t *testing.T) {
	assert.Empty(t, Decline(1))
	assert.Empty(t, Decline(1, "рубль"))
	assert.Equal(t, Decline(1, "рубль", "рубля", "рублей"), "рубль")
	assert.Equal(t, Decline(-4, "рубль", "рубля", "рублей"), "рубля")
	assert.Equal(t, Decline(7, "рубль", "рубля", "рублей"), "рублей")
	assert.Equal(t, Decline(10, "рубль", "рубля", "рублей"), "рублей")
	assert.Equal(t, Decline(-11, "рубль", "рубля", "рублей"), "рублей")
	assert.Equal(t, Decline(17, "рубль", "рубля", "рублей"), "рублей")
	assert.Equal(t, Decline(31, "рубль", "рубля", "рублей"), "рубль")
	assert.Equal(t, Decline(34, "рубль", "рубля", "рублей"), "рубля")
	assert.Equal(t, Decline(-37, "рубль", "рубля", "рублей"), "рублей")
	assert.Equal(t, Decline(101, "рубль", "рубля", "рублей"), "рубль")
	assert.Equal(t, Decline(104, "рубль", "рубля", "рублей"), "рубля")
	assert.Equal(t, Decline(107, "рубль", "рубля", "рублей"), "рублей")
	assert.Empty(t, Decline(1, "dollar"))
	assert.Equal(t, Decline(1, "dollar", "dollars"), "dollar")
	assert.Equal(t, Decline(-4, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(7, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(10, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(-11, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(17, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(31, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(34, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(-37, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(101, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(104, "dollar", "dollars"), "dollars")
	assert.Equal(t, Decline(107, "dollar", "dollars"), "dollars")
}
