package math

import "testing"

// TestSum is a test function for the Sum function.
// It checks if the function returns the correct sum of two integers.
// FILEPATH: math/sum.go
func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
