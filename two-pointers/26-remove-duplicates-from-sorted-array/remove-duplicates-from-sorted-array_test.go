package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		wantK    int
		wantNums []int
	}{
		// Test case 1: Array with distinct elements
		{nums: []int{1, 2, 3, 4, 5}, wantK: 5, wantNums: []int{1, 2, 3, 4, 5}},
		// // Test case 2: Array with one element
		{nums: []int{1}, wantK: 1, wantNums: []int{1}},
		// // Test case 3: Array with all elements the same
		{nums: []int{1, 1, 1, 1}, wantK: 1, wantNums: []int{1}},
		// Test case 4: Increasing array with duplicates
		{nums: []int{1, 1, 2}, wantK: 2, wantNums: []int{1, 2}},
		// Test case 5: Longer array with multiple duplicates
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, wantK: 5, wantNums: []int{0, 1, 2, 3, 4}},
		// Test case 6: Mixed duplicates with gaps
		{nums: []int{1, 2, 2, 3, 4, 4, 5}, wantK: 5, wantNums: []int{1, 2, 3, 4, 5}},
		// Test case 7: Array with no duplicates
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, wantK: 9, wantNums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		// // Test case 8: Alternate duplicates
		{nums: []int{1, 1, 2, 2, 3, 3}, wantK: 3, wantNums: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		numsCopy := make([]int, len(tt.nums))
		copy(numsCopy, tt.nums) // Make a copy to avoid side effects
		gotK := removeDuplicates(numsCopy)
		gotNums := numsCopy[:gotK]

		if gotK != tt.wantK || !reflect.DeepEqual(gotNums, tt.wantNums) {
			t.Errorf("removeDuplicates(%v) = %d, nums = %v; want %d, nums = %v",
				tt.nums, gotK, gotNums, tt.wantK, tt.wantNums)
		}
	}
}
