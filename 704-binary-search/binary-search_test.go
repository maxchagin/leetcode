package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		nums    []int
		target  int
		want    int
		comment string
	}{
		// Test case 1: Example 1 from the prompt
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4, "Example 1: Target 9 exists at index 4"},

		// Test case 2: Example 2 from the prompt
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1, "Example 2: Target 2 does not exist, should return -1"},

		// Test case 3: Target is the first element
		{[]int{1, 2, 3, 4, 5}, 1, 0, "Target is the first element, should return index 0"},

		// Test case 4: Target is the last element
		{[]int{1, 2, 3, 4, 5}, 5, 4, "Target is the last element, should return index 4"},

		// Test case 5: Single element array with target present
		{[]int{8}, 8, 0, "Single element array where target is present, should return index 0"},

		// Test case 6: Single element array with target absent
		{[]int{10}, 7, -1, "Single element array where target is absent, should return -1"},

		// Test case 7: Empty array provided
		{[]int{}, 3, -1, "Empty array, should always return -1"},

		// Test case 8: Large array with target in middle
		{[]int{-10, -5, 0, 5, 10, 15, 20}, 10, 4, "Target 10 exists in a larger array at index 4"},

		// Test case 9: Large array with negative numbers
		{[]int{-20, -15, -10, -5, 0, 5, 10, 20}, -10, 2, "Target -10 found at index 2 amidst negative numbers"},

		// Test case 10: Large negative to positive range, target not present
		{[]int{-50, -30, -20, -10, 0, 10, 20, 30, 50}, 25, -1, "Target 25 is not present, should return -1"},
	}

	for i, c := range cases {
		got := search(c.nums, c.target)
		if got != c.want {
			t.Errorf("Test case %d failed: %s, got %v, want %v", i+1, c.comment, got, c.want)
		}
	}
}
