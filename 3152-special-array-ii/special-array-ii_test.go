package specialarrayi

import (
	"reflect"
	"testing"
)

func TestIsArraySpecial(t *testing.T) {
	cases := []struct {
		nums    []int
		queries [][]int
		want    []bool
		comment string
	}{
		// Test case 1: Example 1 from the problem statement
		{[]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}, []bool{false}, "Example 1: Subarray (3,4,1,2,6) with even pair (2,6)"},

		// Test case 2: Example 2 from the problem statement
		{[]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}, []bool{false, true}, "Example 2: Subarrays (4,3,1) with odd pair (3,1) and (1,6) with different parity"},

		// Test case 3: Single element subarray
		{[]int{5, 8, 3}, [][]int{{1, 1}}, []bool{true}, "Single element subarray, automatically special"},

		// Test case 4: Full array with alternating parity
		{[]int{1, 2, 3, 4, 5}, [][]int{{0, 4}}, []bool{true}, "Alternating parity in full array is special"},

		// Test case 5: Array with all odd numbers
		{[]int{7, 5, 3, 9}, [][]int{{0, 3}}, []bool{false}, "All elements odd, any subarray will be false"},

		// Test case 6: Array with all even numbers
		{[]int{6, 4, 2, 8}, [][]int{{0, 3}}, []bool{false}, "All elements even, the subarray will not be special"},

		// Test case 7: Multiple queries with both possible results
		{[]int{2, 3, 4, 5, 6}, [][]int{{0, 1}, {1, 2}, {0, 4}}, []bool{true, true, true}, "Mixed subarrays with varying parity outcomes"},

		// Test case 8: Complex mix with edge case
		{[]int{1, 2, 1, 2, 1, 2, 1}, [][]int{{0, 6}, {2, 4}}, []bool{true, true}, "Complex mix with subarray having alternating parity"},

		// Test case 9: Edge case query
		{[]int{8}, [][]int{{0, 0}}, []bool{true}, "Single number the full range is special"},

		// Test case 10: Edge case query
		{[]int{1, 1}, [][]int{{0, 1}}, []bool{false}, "Two one"},
	}

	for i, c := range cases {
		got := isArraySpecial(c.nums, c.queries)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Test case %d failed: %s, got %v, want %v", i+1, c.comment, got, c.want)
		}
	}
}
