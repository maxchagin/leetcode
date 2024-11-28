package rotateimage

import (
	"reflect"
	"testing"
)

// TestRotate tests the rotate function with various cases.
func TestRotate(t *testing.T) {
	// Test cases with expected outputs
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		// Basic 3x3 matrix
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		// 4x4 matrix
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		// Single element matrix
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		// 2x2 matrix
		{
			[][]int{{1, 2}, {3, 4}},
			[][]int{{3, 1}, {4, 2}},
		},
		// All elements are the same
		{
			[][]int{{2, 2}, {2, 2}},
			[][]int{{2, 2}, {2, 2}},
		},
		// 5x5 matrix
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			[][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
		// Matrix with negative numbers
		{
			[][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			[][]int{{-7, -4, -1}, {-8, -5, -2}, {-9, -6, -3}},
		},
		// Mixed numbers (positive and negative)
		{
			[][]int{{1, -2, 3}, {-4, 5, -6}, {7, -8, 9}},
			[][]int{{7, -4, 1}, {-8, 5, -2}, {9, -6, 3}},
		},
		// Matrix with zero
		{
			[][]int{{0, 1}, {2, 3}},
			[][]int{{2, 0}, {3, 1}},
		},
		// Large numbers
		{
			[][]int{{100, 200}, {300, 400}},
			[][]int{{300, 100}, {400, 200}},
		},
	}

	for _, test := range tests {
		rotate(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("rotate got %v; want %v", test.input, test.expected)
		}
	}
}
