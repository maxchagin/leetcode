package nextgreaterelementii

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 1}, output: []int{2, -1, 2}},
		{input: []int{1, 2, 3, 4, 3}, output: []int{2, 3, 4, -1, 4}},
		{input: []int{1, 2, 3, 4, 5}, output: []int{2, 3, 4, 5, -1}},
		{input: []int{5, 4, 3, 2, 1}, output: []int{-1, 5, 5, 5, 5}},
		{input: []int{2, 2, 2, 2, 2}, output: []int{-1, -1, -1, -1, -1}},
		{input: []int{}, output: []int{}},
		{input: []int{1}, output: []int{-1}},
		{input: []int{1, 3, 2, 4, 1, 5}, output: []int{3, 4, 4, 5, 5, -1}},
		{input: []int{-1, 0}, output: []int{0, -1}},
		{input: []int{2, 1, 2, 4, 3}, output: []int{4, 2, 4, -1, 4}},
		{input: []int{4, 6, 3, 5, 2, 1}, output: []int{6, -1, 5, 6, 4, 4}},
	}

	for _, test := range tests {
		result := nextGreaterElements(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("For input %v, expected output %v, but got %v", test.input, test.output, result)
		}
	}
}
