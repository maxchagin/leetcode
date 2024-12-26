package searcha2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		// Test case 1: Example from description - target present
		{
			name:     "Basic example with target present",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		// Test case 2: Example from description - target not present
		{
			name:     "Basic example with target absent",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
		// Test case 3: Single row matrix
		{
			name:     "Single row matrix",
			matrix:   [][]int{{1, 3, 5, 7}},
			target:   5,
			expected: true,
		},
		// Test case 4: Single column matrix
		{
			name:     "Single column matrix",
			matrix:   [][]int{{1}, {3}, {5}},
			target:   3,
			expected: true,
		},
		// Test case 5: Target at first position
		{
			name:     "Target at first position",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}},
			target:   1,
			expected: true,
		},
		// Test case 6: Target at last position
		{
			name:     "Target at last position",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}},
			target:   6,
			expected: true,
		},
		// Test case 7: Target smaller than smallest element
		{
			name:     "Target smaller than all elements",
			matrix:   [][]int{{10, 20}, {30, 40}},
			target:   5,
			expected: false,
		},
		// Test case 8: Target larger than largest element
		{
			name:     "Target larger than all elements",
			matrix:   [][]int{{1, 2}, {3, 4}},
			target:   10,
			expected: false,
		},
		// Test case 9: Large matrix with target in middle
		{
			name: "Large matrix with target in middle",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target:   11,
			expected: true,
		},
		// Test case 10: Target between rows
		{
			name:     "Target between rows",
			matrix:   [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}},
			target:   6,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := searchMatrix(test.matrix, test.target); result != test.expected {
				t.Errorf("searchMatrix(%v, %d) = %v, want %v",
					test.matrix, test.target, result, test.expected)
			}
		})
	}
}
