package finalarraystateafterkmultiplicationoperationsi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFinalState(t *testing.T) {
	tests := []struct {
		input      []int
		k          int
		multiplier int
		expected   []int
	}{
		// Test case 1: Basic case with varied numbers and multiple operations
		{input: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2, expected: []int{8, 4, 6, 5, 6}},
		// Test case 2: Basic case with minimal list and more operations than elements
		{input: []int{1, 2}, k: 3, multiplier: 4, expected: []int{16, 8}},
		// Test case 3
		{input: []int{3, 3, 3, 3}, k: 4, multiplier: 2, expected: []int{6, 6, 6, 6}},
		// Test case 4: Minimal length list with one element
		{input: []int{5}, k: 1, multiplier: 3, expected: []int{15}},
		// Test case 5: List with increasing elements, operations within element range
		{input: []int{1, 2, 3, 4, 5}, k: 2, multiplier: 3, expected: []int{3, 6, 3, 4, 5}},
		// Test case 6: Edge case with smallest possible values for all inputs
		{input: []int{1}, k: 1, multiplier: 1, expected: []int{1}},
		// Test case 7: Maximum multiplier effect on all elements being the same
		{input: []int{7, 7, 7, 7}, k: 3, multiplier: 5, expected: []int{35, 35, 35, 7}},
		// Test case 8: Edge case with maximum element value being multiplied
		{input: []int{100, 50, 25, 100}, k: 2, multiplier: 2, expected: []int{100, 100, 50, 100}},
		// Test case 9: Operations exceeding changes needed, testing excess k operations
		{input: []int{2, 2, 2, 2}, k: 10, multiplier: 3, expected: []int{54, 54, 18, 18}},
		// Test case 10: Case where no change after multipliers (multiplier = 1)
		{input: []int{5, 3, 4}, k: 3, multiplier: 1, expected: []int{5, 3, 4}},
	}

	// Iterate over and execute each test case
	for i, test := range tests {
		// Make a copy of the input to preserve its original state when testing
		numsCopy := make([]int, len(test.input))
		copy(numsCopy, test.input)

		// Call the function under test
		result := getFinalState(numsCopy, test.k, test.multiplier)

		// Check if the result matches the expected outcome
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test case %d FAILED:\n", i+1)
			fmt.Printf("Input: %v, k: %d, multiplier: %d\n", test.input, test.k, test.multiplier)
			fmt.Printf("Expected: %v\n", test.expected)
			fmt.Printf("Got: %v\n\n", result)
		} else {
			fmt.Printf("Test case %d passed.\n", i+1)
		}
	}
}
