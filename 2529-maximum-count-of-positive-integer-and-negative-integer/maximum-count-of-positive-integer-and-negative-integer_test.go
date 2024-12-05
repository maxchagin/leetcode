package maximumcountofpositiveintegerandnegativeinteger

import "testing"

func TestMaximumCount(t *testing.T) {
	cases := []struct {
		nums    []int
		want    int
		comment string
	}{
		// Test case 1: Example 1 from the problem statement
		{[]int{-2, -1, -1, 1, 2, 3}, 3, "Example 1: Equal numbers of positive and negative integers"},

		// Test case 2: Example 2 from the problem statement
		{[]int{-3, -2, -1, 0, 0, 1, 2}, 3, "Example 2: More negative integers than positive integers"},

		// Test case 3: Example 3 from the problem statement
		{[]int{5, 20, 66, 1314}, 4, "Example 3: All positive integers, no negative integers"},

		// Test case 4: No positive integers
		{[]int{-5, -4, -3, -2, -1}, 5, "All negative integers, no positives"},

		// Test case 5: No negative integers
		{[]int{1, 2, 3, 4, 5}, 5, "All positive integers, no negatives"},

		// Test case 6: Mixed positive and negative integers with gaps
		{[]int{-10, -5, 0, 1, 2, 3}, 3, "Mixed values with zero, fewer negatives"},

		// Test case 7: Only zero values
		{[]int{0, 0, 0, 0}, 0, "Only zeros, zero negative or positive integers"},

		// Test case 8: Single negative integer
		{[]int{-1}, 1, "Single negative integer, should be counted as one"},

		// Test case 9: Single positive integer
		{[]int{1}, 1, "Single positive integer, should be counted as one"},

		// Test case 10: Larger input with more positive integers
		{[]int{-3, -2, -1, 0, 1, 2, 3, 4, 5}, 5, "More positives than negatives"},
	}

	for i, c := range cases {
		got := maximumCount(c.nums)
		if got != c.want {
			t.Errorf("Test case %d failed: %s, got %v, want %v", i+1, c.comment, got, c.want)
		}
	}
}
