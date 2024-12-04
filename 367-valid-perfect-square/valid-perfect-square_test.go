package validperfectsquare

import "testing"

// TestIsPerfectSquare tests the isPerfectSquare function.
func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		// Case 1: The input is a perfect square.
		{
			input:    16,
			expected: true, // 4 * 4 = 16
		},
		// Case 2: The input is not a perfect square.
		{
			input:    14,
			expected: false, // Between 3 * 3 = 9 and 4 * 4 = 16
		},
		// Case 3: The smallest perfect square.
		{
			input:    1,
			expected: true, // 1 * 1 = 1
		},
		// Case 4: Larger perfect square.
		{
			input:    100,
			expected: true, // 10 * 10 = 100
		},
		// Case 5: Larger non-perfect square.
		{
			input:    99,
			expected: false, // Between 9 * 9 = 81 and 10 * 10 = 100
		},
		// Case 6: Maximum possible input (edge case for the unsigned int constraint).
		{
			input:    2147483647,
			expected: false, // 2147483647 is not a perfect square
		},
		// Case 7: Largest perfect square less than the 32-bit limit.
		{
			input:    2147395600,
			expected: true, // 46340 * 46340 = 2147395600
		},
		// Case 8: Perfect square of a small integer.
		{
			input:    81,
			expected: true, // 9 * 9 = 81
		},
		// Case 9: Non-perfect square with large difference to next perfect square.
		{
			input:    50,
			expected: false, // Between 7 * 7 = 49 and 8 * 8 = 64
		},
		// Case 10: Even number that is not a perfect square.
		{
			input:    18,
			expected: false, // Between 4 * 4 = 16 and 5 * 5 = 25
		},
	}

	for _, test := range tests {
		output := isPerfectSquare(test.input)
		if output != test.expected {
			t.Errorf("For input %d, expected %v, but got %v", test.input, test.expected, output)
		}
	}
}
