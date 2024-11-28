package beautifultowersi

import "testing"

func TestMaximumSumOfHeights(t *testing.T) {
	tests := []struct {
		heights     []int
		expected    int64
		description string
	}{
		// Basic test case where the peak is at the beginning
		{[]int{5, 3, 4, 1, 1}, 13, "Peak at the beginning"},

		// Test case with peak in the middle
		{[]int{6, 5, 3, 9, 2, 7}, 22, "Peak in the middle"},

		// Test case with consecutive equal peak heights
		{[]int{3, 2, 5, 5, 2, 3}, 18, "Consecutive equal peak heights"},

		// Test case with all towers of the same height
		{[]int{5, 5, 5, 5, 5}, 25, "All towers same height"},

		// Test case with strictly increasing then decreasing heights
		{[]int{1, 2, 3, 4, 3, 2, 1}, 16, "Perfect mountain shape"},

		// Test case with only one tower
		{[]int{10}, 10, "Single tower"},

		// Test case with heights in decreasing order only
		{[]int{5, 4, 3, 2, 1}, 15, "Strictly decreasing heights"},

		// Test case with heights in increasing order only
		{[]int{1, 2, 3, 4, 5}, 15, "Strictly increasing heights"},

		// Edge case with an empty array
		{[]int{}, 0, "Empty array"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := maximumSumOfHeights(test.heights)
			if result != test.expected {
				t.Errorf("For heights %v, expected maximum sum %d, but got %d", test.heights, test.expected, result)
			}
		})
	}
}
