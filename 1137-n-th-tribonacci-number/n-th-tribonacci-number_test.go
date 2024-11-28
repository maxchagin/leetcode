package nthtribonaccinumber

import "testing"

// TestTribonacci tests the tribonacci function with various cases.
func TestTribonacci(t *testing.T) {
	// Test cases with expected outputs
	tests := []struct {
		input    int
		expected int
	}{
		// Basic case: T0
		{0, 0},
		// Simple base case: T1
		{1, 1},
		// Simple base case: T2
		{2, 1},
		// Smallest n where calculation is needed
		{3, 2},
		// Calculated with Tn = Tn-1 + Tn-2 + Tn-3
		{4, 4},
		// Example from problem statement
		{25, 1389537},
		// Intermediate value test: T10
		{10, 149},
		// Test for T5
		{5, 7},
		// Larger value check for correctness
		{15, 3136},
		// Another value to ensure consistency
		{20, 66012},
	}

	for _, test := range tests {
		result := tribonacci(test.input)
		if result != test.expected {
			t.Errorf("tribonacci(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
