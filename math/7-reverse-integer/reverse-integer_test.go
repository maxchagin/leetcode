package reverseinteger

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		// Test case 1: Simple positive number
		{123, 321},
		// Test case 2: Simple negative number
		{-123, -321},
		// // Test case 3: Positive number ending with zero
		{120, 21},
		// // Test case 4: Positive number ending with zero
		{101, 101},
		// Test case 5: Zero input
		{0, 0},
		// Test case 6: Number that leads to overflow after reversal
		{1534236469, 0}, // Reversing this number causes overflow
		// Test case 7: Negative number that leads to overflow after reversal
		{-1563847412, 0}, // Reversal causes negative overflow
		// Test case 8: Number with multiple zeros at the end
		{1000, 1},
		// Test case 9: Large negative number within 32-bit range
		{-2147483412, -2143847412},
		// Test case 10: Number whose reverse is zero due to leading zeros
		{-10, -1},
		// Test case 11: Edge case of largest negative 32-bit integer
		{-2147483648, 0}, // Reversing causes overflow
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reverse(%d)", test.input), func(t *testing.T) {
			result := reverse(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
