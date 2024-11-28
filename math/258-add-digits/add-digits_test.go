package adddigits

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		// Test case 1: Simple number with two digits
		{38, 2},
		// Test case 2: Single digit number
		{0, 0},
		// Test case 3: Check for a number that reduces to itself
		{9, 9},
		// Test case 4: Larger number with multiple additions
		{123, 6}, // 1 + 2 + 3 = 6
		// Test case 5: Number divisible by 9 (non-zero)
		{18, 9}, // 1 + 8 = 9
		// Test case 6: Large multi-digit number
		{987654321, 9}, // Sum of digits is 45, which reduces to 9
		// Test case 7: Various digit repetition
		{111111, 6}, // Sum of digits is 6x1 = 6
		// Test case 8: Reduction involving multiple steps
		{49, 4}, // 4 + 9 = 13, then 1 + 3 = 4
		// Test case 9: Power of 10 minus one
		{99, 9}, // 9 + 9 = 18, then 1 + 8 = 9
		// Test case 10: Maximum input constraint
		{2147483647, 1}, // Sum of digits is 46, then reduces to 1
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("addDigits(%d)", test.input), func(t *testing.T) {
			result := addDigits(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
