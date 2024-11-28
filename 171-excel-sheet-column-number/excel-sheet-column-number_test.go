package excelsheetcolumnnumber

import (
	"fmt"
	"testing"
)

// Test case structure for the titleToNumber function
type testCase struct {
	columnTitle string
	expected    int
}

// TestTitleToNumber tests the titleToNumber function with various test cases.
func TestTitleToNumber(t *testing.T) {
	tests := []testCase{
		// Test case 1: Single letter, simplest possible case
		{columnTitle: "A", expected: 1},

		// Test case 2: End of single-letter range
		{columnTitle: "Z", expected: 26},

		// Test case 3: Transition from single to double letters
		{columnTitle: "AA", expected: 27},

		// Test case 4: Standard double-letter test
		{columnTitle: "AB", expected: 28},

		// Test case 5: End of double-letter range
		{columnTitle: "ZZ", expected: 702},

		// Test case 6: Start of triple-letter columns
		{columnTitle: "AAA", expected: 703},

		// Test case 7: Test with mixed letters
		{columnTitle: "ABC", expected: 731},

		// Test case 8: High-value computation
		{columnTitle: "XYZ", expected: 16900},

		// Test case 9: Another mixed letters test
		{columnTitle: "AZY", expected: 1377},

		// Test case 10: Boundary condition with a large input
		{columnTitle: "AAAA", expected: 18279}, // 26^3 = 17576 + 26^1 + 26^0
	}

	for _, tt := range tests {
		// Use fmt.Sprintf to format the sub-test name
		t.Run(fmt.Sprintf("columnTitle_%s", tt.columnTitle), func(t *testing.T) {
			result := titleToNumber(tt.columnTitle)
			if result != tt.expected {
				t.Errorf("titleToNumber(%s) = %v, expected %v", tt.columnTitle, result, tt.expected)
			}
		})
	}
}
