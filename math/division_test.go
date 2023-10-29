package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDivision is a test function for the Division function using testify.
func TestDivision(t *testing.T) {
	assert.Equal(t, 2.0, Division(4, 2), "4 / 2 should be 2.0")

	assert.NotEqual(t, 3.0, Division(5, 2), "5 / 2 should not be 3.0")
}
