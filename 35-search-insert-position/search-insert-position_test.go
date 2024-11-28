package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6, 8, 10}, 9, 5},
	}

	for _, test := range tests {
		result := searchInsert(test.nums, test.target)
		if result != test.expected {
			t.Errorf("searchInsert(%v, %d) = %d; expected %d", test.nums, test.target, result, test.expected)
		}
	}
}
