package maxchunkstomakesorted

import (
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		// Test case 1: Single chunk
		{arr: []int{4, 3, 2, 1, 0}, expected: 1},
		// Test case 2: Multiple chunks
		{arr: []int{1, 0, 2, 3, 4}, expected: 4},
		// Test case 3: Already sorted array
		{arr: []int{0, 1, 2, 3, 4}, expected: 5},
		// Test case 4: Two chunks
		{arr: []int{2, 0, 1, 3, 4}, expected: 3},
		// Test case 5: Single element array
		{arr: []int{0}, expected: 1},
		// Test case 6: Two elements swapped
		{arr: []int{1, 0}, expected: 1},
		// Test case 7: Two chunks
		{arr: []int{1, 2, 0, 4, 3}, expected: 2},
		// Test case 8: Three chunks
		{arr: []int{1, 0, 3, 2, 4}, expected: 3},
		// Test case 9: No chunks possible
		{arr: []int{4, 3, 2, 0, 1}, expected: 1},
		// Test case 10: Random order
		{arr: []int{3, 2, 4, 1, 0}, expected: 1},
	}

	for _, test := range tests {
		if result := maxChunksToSorted(test.arr); result != test.expected {
			t.Errorf("For arr %v, expected %d, but got %d", test.arr, test.expected, result)
		}
	}
}
