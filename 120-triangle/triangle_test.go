package triangle

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		input       [][]int
		expected    int
		description string
	}{
		// Test Case 1: Basic triangle with distinct numbers
		{input: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, expected: 11, description: "Basic example with multiple levels."},

		// Test Case 2: Single element triangle
		{input: [][]int{{-10}}, expected: -10, description: "Single element, minimum path is the element itself."},

		// Test Case 3: Two levels triangle
		{input: [][]int{{1}, {2, 3}}, expected: 3, description: "Two levels, choosing the smaller number in the second level."},

		// Test Case 4: Triangle with negative numbers included
		{input: [][]int{{-1}, {2, 3}, {1, -1, -3}}, expected: -1, description: "Triangle with negative numbers, path should consider negatives."},

		// Test Case 5: All elements are zero
		{input: [][]int{{0}, {0, 0}, {0, 0, 0}}, expected: 0, description: "All elements zero, resulting path sum should be zero."},

		// Test Case 6: Larger triangle with increasing elements
		{input: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}, {1, 1, 1, 1, 1}}, expected: 12, description: "Larger triangle, should correctly sum minimum path."},

		// Test Case 7: Path choosing on edges
		{input: [][]int{{5}, {9, 6}, {4, 6, 8}, {0, 7, 1, 5}}, expected: 18, description: "Path that involves choosing elements from corners."},

		// Test Case 8: All negative triangle
		{input: [][]int{{-1}, {-2, -3}, {-4, -5, -6}}, expected: -10, description: "All elements are negative, minimum path is addition of all elements."},

		// Test Case 9: One path triangle
		{input: [][]int{{10}, {9, 8}, {7, 6, 5}, {4, 3, 2, 1}}, expected: 24, description: "Single path from top to bottom."},

		// Test Case 10: Random numbers triangle
		{input: [][]int{{4}, {1, 2}, {7, 3, 6}, {8, 5, 9, 3}, {3, 7, 2, 5, 1}}, expected: 15, description: "Random numbers with mixed possibilities."},

		// Test Case 11: Leetcode testcase
		{input: [][]int{{-1}, {-2, -3}}, expected: -4, description: "Leetcode testcase 11."},

		// Test Case 12: Leetcode testcase
		{input: [][]int{{-1}, {2, 3}, {1, -1, -3}}, expected: -1, description: "Leetcode testcase 12."},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := minimumTotal(test.input)
			if output != test.expected {
				t.Errorf("Test case '%s' failed: got %d, expected %d", test.description, output, test.expected)
			}
		})
	}
}
