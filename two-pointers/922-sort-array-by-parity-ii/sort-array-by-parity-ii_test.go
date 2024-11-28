package sortarraybyparityii

import (
	"reflect"
	"testing"
)

func TestSortArrayByParityII(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		// Test case 1: Simple case with two pairs of even and odd numbers
		{nums: []int{4, 2, 5, 7}, want: []int{4, 5, 2, 7}}, // One possible output
		// Test case 2: Only two elements, one even and one odd
		{nums: []int{2, 3}, want: []int{2, 3}},
		// Test case 5: Larger array with a balanced number of even and odd numbers
		{nums: []int{9, 8, 7, 6, 5, 4, 3, 2}, want: []int{8, 9, 6, 7, 4, 5, 2, 3}}, // One possible output
		// Test case 6: Alternating odd and even starting with odd
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: []int{2, 1, 4, 3, 6, 5, 8, 7}}, // One possible output
		// Test case 7: Alternating even and odd starting with even
		{nums: []int{2, 1, 4, 3, 6, 5, 8, 7}, want: []int{2, 1, 4, 3, 6, 5, 8, 7}},
		// Test case 8: Array with minimal differences
		{nums: []int{2, 4, 1, 3}, want: []int{2, 1, 4, 3}},
		// Test case 9: Reverse order input array
		{nums: []int{9, 7, 5, 3, 8, 6, 4, 2}, want: []int{8, 9, 6, 7, 4, 5, 2, 3}}, // One possible output
		// Test case 10: Even numbers are larger
		{nums: []int{10, 3, 14, 7, 8, 1}, want: []int{10, 3, 14, 7, 8, 1}},
		{nums: []int{3, 0, 4, 0, 2, 1, 3, 1, 3, 4}, want: []int{0, 3, 4, 3, 2, 1, 0, 1, 4, 3}},
		{nums: []int{1, 4, 4, 3, 0, 3}, want: []int{4, 1, 4, 3, 0, 3}},
	}

	for _, tt := range tests {
		tmpNums := make([]int, len(tt.nums))
		copy(tmpNums, tt.nums)
		got := sortArrayByParityII(tmpNums)
		if !isValidParityPosition(got) || !containsSameElements(got, tt.want) {
			t.Errorf("sortArrayByParityII(%v) = %v; want valid parity arrangement like %v", tt.nums, got, tt.want)
		}
	}
}

// Helper function to check if output has valid parity positions
func isValidParityPosition(nums []int) bool {
	for i, num := range nums {
		if i%2 != num%2 {
			return false
		}
	}
	return true
}

// Helper to check if the sorted result contains the same elements as wanted
func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create a map to count elements
	countA := make(map[int]int)
	countB := make(map[int]int)

	for _, num := range a {
		countA[num]++
	}

	for _, num := range b {
		countB[num]++
	}

	return reflect.DeepEqual(countA, countB)
}
