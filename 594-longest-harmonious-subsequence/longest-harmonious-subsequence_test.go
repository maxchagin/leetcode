package longestharmonioussubsequence

import "testing"

func TestFindLHS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Standard test cases
		{
			name:     "Example 1",
			nums:     []int{1, 3, 2, 2, 5, 2, 3, 7},
			expected: 5, // [3,2,2,2,3]
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			expected: 2, // [1,2], [2,3], or [3,4]
		},
		{
			name:     "Example 3 (No harmonious subsequence)",
			nums:     []int{1, 1, 1, 1},
			expected: 0, // All elements are the same
		},

		// Edge cases
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 0, // Only one element, no possible subsequence
		},
		{
			name:     "Two elements with diff 1",
			nums:     []int{5, 6},
			expected: 2, // Valid harmonious subsequence
		},
		{
			name:     "Two elements with diff > 1",
			nums:     []int{5, 7},
			expected: 0, // Diff is 2 (invalid)
		},
		{
			name:     "Negative numbers with valid diff",
			nums:     []int{-1, -2, -2, -3},
			expected: 3, // [-2, -2, -1] or [-2, -2, -3]
		},

		// Larger input cases
		{
			name:     "Repeated pairs with diff 1",
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 2},
			expected: 7, // [1, 2, 1, 2, 1, 2, 2]
		},
		{
			name:     "All elements same except one",
			nums:     []int{5, 5, 5, 5, 6},
			expected: 5, // [5,5,5,5,6]
		},
		{
			name:     "Leetcode case 1",
			nums:     []int{1, 2, 1, 3, 0, 0, 2, 2, 1, 3, 3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.nums); got != tt.expected {
				t.Errorf("findLHS() = %v, want %v", got, tt.expected)
			}
		})
	}
}
