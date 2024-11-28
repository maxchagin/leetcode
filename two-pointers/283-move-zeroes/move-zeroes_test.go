package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		// Test case 1: Mixed zeros and non-zeros
		{nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		// Test case 2: Single zero element
		{nums: []int{0}, want: []int{0}},
		// Test case 3: No zero elements
		{nums: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		// Test case 4: All elements are zeros
		{nums: []int{0, 0, 0, 0}, want: []int{0, 0, 0, 0}},
		// Test case 5: Zeros at the start
		{nums: []int{0, 0, 1, 2, 3}, want: []int{1, 2, 3, 0, 0}},
		// Test case 6: Zeros at the end
		{nums: []int{1, 2, 3, 0, 0}, want: []int{1, 2, 3, 0, 0}},
		// Test case 7: Alternating zeros and non-zeros
		{nums: []int{0, 1, 0, 2, 0, 3}, want: []int{1, 2, 3, 0, 0, 0}},
		// Test case 8: Single non-zero, multiple zeros
		{nums: []int{0, 0, 0, 1}, want: []int{1, 0, 0, 0}},
		// Test case 9: Large array with random zeros
		{nums: []int{1, 0, 2, 0, 0, 3, 4, 0, 5, 0}, want: []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}},
		// Test case 10: Non-zero at end
		{nums: []int{0, 0, 0, 0, 1}, want: []int{1, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		got := make([]int, len(tt.nums))
		copy(got, tt.nums)
		moveZeroes(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("moveZeroes(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
