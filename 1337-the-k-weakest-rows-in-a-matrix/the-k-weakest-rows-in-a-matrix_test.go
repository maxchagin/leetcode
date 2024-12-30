package thekweakestrowsinamatrix

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	tests := []struct {
		mat      [][]int
		k        int
		expected []int
	}{
		{
			// Test case 1: Example case from the prompt
			mat:      [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}},
			k:        3,
			expected: []int{2, 0, 3},
		},
		{
			// Test case 2: Another example case from the prompt
			mat:      [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			k:        2,
			expected: []int{0, 2},
		},
		{
			// Test case 3: All rows have the same number of soldiers
			mat:      [][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 0}},
			k:        2,
			expected: []int{0, 1},
		},
		{
			// Test case 4: Rows with no soldiers
			mat:      [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			k:        2,
			expected: []int{0, 1},
		},
		{
			// Test case 5: Rows with all soldiers
			mat:      [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			k:        2,
			expected: []int{0, 1},
		},
		{
			// Test case 6: Mixed rows
			mat:      [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}},
			k:        2,
			expected: []int{0, 1},
		},
		{
			// Test case 7: Single row matrix
			mat:      [][]int{{1, 1, 1, 1}},
			k:        1,
			expected: []int{0},
		},
		{
			// Test case 8: Single column matrix
			mat:      [][]int{{1}, {0}, {1}},
			k:        2,
			expected: []int{1, 0},
		},
		{
			// Test case 9: Large matrix with varying soldiers
			mat:      [][]int{{1, 1, 1, 0}, {1, 1, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 1}},
			k:        3,
			expected: []int{2, 1, 0},
		},
		{
			// Test case 10
			mat:      [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			k:        2,
			expected: []int{0, 2},
		},
		{
			// Test case 11
			mat:      [][]int{{1, 1}, {1, 0}, {1, 0}, {1, 1}, {0, 0}, {1, 1}},
			k:        1,
			expected: []int{4},
		},
	}

	for _, test := range tests {
		result := kWeakestRows(test.mat, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("kWeakestRows(%v, %d) = %v; expected %v", test.mat, test.k, result, test.expected)
		}
	}
}
