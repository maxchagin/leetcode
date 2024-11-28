package pascalstriangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Define the test cases using a struct
	tests := []struct {
		input    int     // Number of rows to generate in Pascal's Triangle
		expected [][]int // Expected 2D slice output for numRows
	}{
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}}, // Example 1 from the problem statement
		{8, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
			{1, 5, 10, 10, 5, 1},
			{1, 6, 15, 20, 15, 6, 1},
			{1, 7, 21, 35, 35, 21, 7, 1},
		}}, // Example 2 from the problem statement

		{1, [][]int{
			{1},
		}}, // Example 3 from the problem statement

		{0, [][]int{}}, // Edge case: zero rows should output an empty slice

		{2, [][]int{
			{1},
			{1, 1},
		}}, // Simple case: two rows

		{3, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
		}}, // Simple case: three rows
	}

	// Iterate over each test case
	for i, test := range tests {
		result := generate(test.input) // Call the function with the test input

		// Verify if the result matches the expected output
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: For input %d, expected %v, but got %v",
				i, test.input, test.expected, result)
		}
	}
}
