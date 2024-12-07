package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		// Test with positive numbers
		{input: []int{1, 2, 3, 4}, output: []int{24, 12, 8, 6}},

		// Test with negative numbers
		{input: []int{-1, -2, -3, -4}, output: []int{-24, -12, -8, -6}},

		// Test with zero in the array
		{input: []int{-1, 1, 0, -3, 3}, output: []int{0, 0, 9, 0, 0}},

		// Test with multiple zeros in the array
		{input: []int{0, 0, 3, 4}, output: []int{0, 0, 0, 0}},

		// Test with one zero but multiple other elements
		{input: []int{4, 0, 2, 10}, output: []int{0, 80, 0, 0}},

		// Test with two elements in the array
		{input: []int{5, 7}, output: []int{7, 5}},

		// Test where all elements are the same and positive
		{input: []int{3, 3, 3, 3}, output: []int{27, 27, 27, 27}},

		// Test where all elements are the same and negative
		{input: []int{-2, -2, -2, -2}, output: []int{-8, -8, -8, -8}},

		// Test with the largest constraint value
		{input: []int{30, 3, 5, 2}, output: []int{30, 300, 180, 450}},

		// Test when input is at its minimum length boundary
		{input: []int{1, 2}, output: []int{2, 1}},

		// Test when two zero
		{input: []int{2, 3, 0, 0}, output: []int{0, 0, 0, 0}},

		// Test
		{input: []int{0, 2, 3, 4}, output: []int{24, 0, 0, 0}},
	}

	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		result := productExceptSelf(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", input, test.output, result)
		}
	}
}
