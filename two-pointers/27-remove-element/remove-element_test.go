package removeelement

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		wantK    int
		wantNums []int
	}{
		// Test case 1: Elements to remove are at the start and end
		{nums: []int{3, 2, 2, 3}, val: 3, wantK: 2, wantNums: []int{2, 2}},
		// Test case 2: Elements to remove are in the middle
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, wantK: 5, wantNums: []int{0, 1, 3, 0, 4}},
		// Test case 3: No elements to remove
		{nums: []int{1, 2, 3, 4, 5}, val: 6, wantK: 5, wantNums: []int{1, 2, 3, 4, 5}},
		// Test case 4: All elements to remove
		{nums: []int{2, 2, 2, 2}, val: 2, wantK: 0, wantNums: []int{}},
		// Test case 5: Alternate removal
		{nums: []int{1, 2, 1, 2, 1, 2}, val: 2, wantK: 3, wantNums: []int{1, 1, 1}},
		// Test case 6: Single element, not equal to val
		{nums: []int{5}, val: 2, wantK: 1, wantNums: []int{5}},
		// Test case 7: Single element, equal to val
		{nums: []int{2}, val: 2, wantK: 0, wantNums: []int{}},
		// Test case 8: Mixed values to remove in a larger array
		{nums: []int{2, 3, 4, 5, 2, 6, 2, 7}, val: 2, wantK: 5, wantNums: []int{3, 4, 5, 6, 7}},
		// Test case 9: Empty array
		{nums: []int{}, val: 1, wantK: 0, wantNums: []int{}},
		// Test case 10: Duplicates of val inter-mixed
		{nums: []int{2, 1, 2, 1, 2, 1, 2, 1}, val: 2, wantK: 4, wantNums: []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		numsCopy := make([]int, len(tt.nums))
		copy(numsCopy, tt.nums) // Make a copy to avoid side effects
		gotK := removeElement(numsCopy, tt.val)

		// Sort and compare the first k elements only
		gotNums := numsCopy[:gotK]
		sort.Ints(gotNums) // Sort to ensure comparison doesn't fail due to order
		sort.Ints(tt.wantNums)

		if gotK != tt.wantK || !reflect.DeepEqual(gotNums, tt.wantNums) {
			t.Errorf("removeElement(%v, %d) = %d, nums = %v; want %d, nums = %v",
				tt.nums, tt.val, gotK, gotNums, tt.wantK, tt.wantNums)
		}
	}
}
