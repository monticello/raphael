package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSession(t *testing.T) {
	id := &UserSession{}

	authenticated1 := id.Authenticate("equal", "different")
	assert.False(t, authenticated1, "It should be false")

	authenticated2 := id.Authenticate("equal", "equal")
	assert.True(t, authenticated2, "It should be true")
}
