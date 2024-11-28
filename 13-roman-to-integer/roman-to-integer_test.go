package romantointeger

import "testing"

// TestRomanToInt tests the romanToInt function with various cases.
func TestRomanToInt(t *testing.T) {
	// Test cases with expected outputs
	tests := []struct {
		input    string
		expected int
	}{
		// Basic case: Single character
		{"I", 1},

		// Basic addition without subtraction
		{"III", 3},

		// Single subtraction case with two characters
		{"IV", 4},

		// Mixed with addition and subtraction
		{"IX", 9},

		// Larger single numeral
		{"LVIII", 58},

		// Mixed numerals including various combinations
		{"MCMXCIV", 1994},

		// Checking consecutive numerals
		{"MMM", 3000},

		// Complex subtraction case
		{"CDXLIV", 444},

		// Only large numerals
		{"MMXXII", 2022},

		// Adding and subtraction around hundreds and thousands
		{"MDCCLXXVI", 1776},
	}

	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.expected {
			t.Errorf("romanToInt(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
