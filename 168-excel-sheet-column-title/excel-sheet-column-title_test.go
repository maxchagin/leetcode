package excelsheetcolumntitle

import (
	"fmt"
	"testing"
)

// Test case structure for the convertToTitle function
type testCase struct {
	columnNumber int
	expected     string
}

// TestConvertToTitle tests the convertToTitle function with various test cases.
func TestConvertToTitle(t *testing.T) {
	tests := []testCase{
		// Test case 1: Smallest number represents the first letter
		{columnNumber: 1, expected: "A"},

		// Test case 2: Transition from single letter to double letters
		{columnNumber: 26, expected: "Z"},

		// Test case 3: First double-letter case
		{columnNumber: 27, expected: "AA"},

		// Test case 4: Typical case with two letters
		{columnNumber: 28, expected: "AB"},

		// Test case 5: Transition at higher numbers
		{columnNumber: 52, expected: "AZ"},

		// Test case 6: Check deeper double-letter structure
		{columnNumber: 701, expected: "ZY"},

		// Test case 7: First of triple letters
		{columnNumber: 703, expected: "AAA"},

		// Test case 8: Another triple-letter test case
		{columnNumber: 731, expected: "ABC"},

		// Test case 9: Boundary case with max two-letter combo
		{columnNumber: 702, expected: "ZZ"},

		// Test case 10: Larger triple-letter case
		{columnNumber: 1377, expected: "AZY"},
	}

	for _, tt := range tests {
		// Use fmt.Sprintf for proper integer to string conversion
		t.Run(fmt.Sprintf("columnNumber_%d", tt.columnNumber), func(t *testing.T) {
			result := convertToTitle(tt.columnNumber)
			if result != tt.expected {
				t.Errorf("convertToTitle(%d) = %s, expected %v", tt.columnNumber, result, tt.expected)
			}
		})
	}
}
