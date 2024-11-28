package sortarraybyparity

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		// Test case 1: Simple case with both even and odd numbers
		{nums: []int{3, 1, 2, 4}, want: []int{2, 4, 3, 1}}, // One possible output
		// Test case 2: Single element, even number
		{nums: []int{0}, want: []int{0}},
		// Test case 3: All even numbers
		{nums: []int{2, 4, 6, 8}, want: []int{2, 4, 6, 8}},
		// Test case 4: All odd numbers
		{nums: []int{1, 3, 5, 7}, want: []int{1, 3, 5, 7}},
		// Test case 5: Mixed numbers with no particular order
		{nums: []int{7, 3, 8, 5, 2, 1, 4}, want: []int{8, 2, 4, 7, 3, 5, 1}}, // One possible output
		// Test case 6: Array with two elements
		{nums: []int{1, 2}, want: []int{2, 1}},
		// Test case 7: Single element, odd number
		{nums: []int{1}, want: []int{1}},
		// Test case 8: Alternating even and odd numbers
		{nums: []int{2, 5, 4, 7, 6, 9}, want: []int{2, 4, 6, 5, 7, 9}}, // One possible output
		// Test case 9: Empty array
		{nums: []int{}, want: []int{}},
		// Test case 10: Large array with repeating numbers
		{nums: []int{1, 1, 2, 2, 1, 2}, want: []int{2, 2, 2, 1, 1, 1}}, // One possible output
		{nums: []int{0, 1, 2}, want: []int{0, 2, 1}},
	}

	for _, tt := range tests {
		got := sortArrayByParity(tt.nums)
		if !isValidSort(got, tt.want) {
			t.Errorf("sortArrayByParity(%v) = %v; expected elements %v", tt.nums, got, tt.want)
		}
	}
}

// Helper function to check valid parity order
func isValidSort(sorted []int, want []int) bool {
	// Check length first
	if len(sorted) != len(want) {
		return false
	}

	// Check if sorted array has been shuffled in a valid way
	for i := 0; i < len(sorted); i++ {
		if sorted[i]%2 != 0 {
			for j := i; j < len(sorted); j++ {
				if sorted[j]%2 == 0 {
					return false
				}
			}
			break
		}
	}

	// Check if both arrays cover the same set (order can be different)
	return reflect.DeepEqual(countElements(sorted), countElements(want))
}

// Helper to count elements for deep comparison
func countElements(arr []int) map[int]int {
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}
	return count
}
