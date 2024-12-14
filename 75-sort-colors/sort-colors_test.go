package sortcolors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		// Test case 1: General case with mixed colors
		{input: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		// Test case 2: General case with all three colors
		{input: []int{2, 0, 1}, expected: []int{0, 1, 2}},
		// Test case 3: Already sorted array
		{input: []int{0, 0, 1, 1, 2, 2}, expected: []int{0, 0, 1, 1, 2, 2}},
		// Test case 4: Reverse sorted array
		{input: []int{2, 2, 1, 1, 0, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		// Test case 5: Array with only zeros
		{input: []int{0, 0, 0}, expected: []int{0, 0, 0}},
		// Test case 6: Array with only ones
		{input: []int{1, 1, 1}, expected: []int{1, 1, 1}},
		// Test case 7: Array with only twos
		{input: []int{2, 2, 2}, expected: []int{2, 2, 2}},
		// Test case 8: Array with two colors (zeros and ones)
		{input: []int{0, 1, 0, 1}, expected: []int{0, 0, 1, 1}},
		// Test case 9: Single element array
		{input: []int{1}, expected: []int{1}},
		// Test case 10: Empty array
		{input: []int{}, expected: []int{}},
		// Test case 11
		{input: []int{0, 2}, expected: []int{0, 2}},
		// Test case 12
		{input: []int{2, 0}, expected: []int{0, 2}},
	}

	// Iterate over test cases
	for i, test := range tests {
		// Make a copy of the input to avoid modifying the original test data
		nums := make([]int, len(test.input))
		copy(nums, test.input)

		// Call the function under test
		sortColors(nums)

		// Check if the output matches the expected result
		if !reflect.DeepEqual(nums, test.expected) {
			fmt.Printf("Test case %d FAILED:\n", i+1)
			fmt.Printf("Input: %v\n", test.input)
			fmt.Printf("Expected: %v\n", test.expected)
			fmt.Printf("Got: %v\n\n", nums)
		} else {
			fmt.Printf("Test case %d passed.\n", i+1)
		}
	}
}
