package findtriangularsumofanarray

import "testing"

func TestTriangularSum(t *testing.T) {
	tests := []struct {
		input       []int
		expected    int
		description string
	}{
		// Test Case 1: Basic case
		{input: []int{1, 2, 3, 4, 5}, expected: 8, description: "Basic example with multiple elements."},

		// Test Case 2: Single element in input
		{input: []int{5}, expected: 5, description: "Single element, result is the element itself."},

		// Test Case 3: All elements are the same
		{input: []int{1, 1, 1, 1, 1}, expected: 6, description: "All elements same, should result in modulo additions."},

		// Test Case 4: Descending order elements
		{input: []int{9, 8, 7, 6, 5}, expected: 2, description: "Descending elements, to check processing."},

		// Test Case 5: Ascending order elements
		{input: []int{0, 1, 2, 3, 4, 5, 9}, expected: 5, description: "Ascending elements and check for carrying forward values."},

		// Test Case 6: Handling overflow with max digits
		{input: []int{9, 9, 9, 9}, expected: 2, description: "All elements are 9, checks modulo overflow handling."},

		// Test Case 7: Alternating zeros and nines
		{input: []int{9, 0, 9, 0, 9, 0}, expected: 4, description: "Checking alternation handling and modular results."},

		// Test Case 8: Random elements
		{input: []int{4, 7, 3, 2, 5}, expected: 3, description: "Random unordered elements in the list."},

		// Test Case 9: Two-element array
		{input: []int{3, 6}, expected: 9, description: "Two elements, simple addition and modulo."},

		// Test Case 10: Large input case
		{input: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, expected: 4, description: "Larger input with repetitive elements."},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := triangularSum(test.input)
			if output != test.expected {
				t.Errorf("Test case '%s' failed: got %d, expected %d", test.description, output, test.expected)
			}
		})
	}
}
