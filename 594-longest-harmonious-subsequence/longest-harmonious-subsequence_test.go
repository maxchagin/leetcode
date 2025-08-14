package longestharmonioussubsequence

import "testing"

func TestFindLHS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Example 1: classic example with answer 5
		{
			name:     "Example 1",
			nums:     []int{1, 3, 2, 2, 5, 2, 3, 7},
			expected: 5, // [3,2,2,2,3]
		},
		// Example 2: consecutive numbers
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			expected: 2, // [1,2] or [2,3] etc.
		},
		// Example 3: all same numbers, no harmonious subsequence
		{
			name:     "Example 3",
			nums:     []int{1, 1, 1, 1},
			expected: 0,
		},
		// Edge case: minimum size array, no harmonious subsequence
		{
			name:     "Small array without harmonious",
			nums:     []int{5},
			expected: 0,
		},
		// Two numbers with difference 1
		{
			name:     "Two consecutive numbers",
			nums:     []int{2, 3},
			expected: 2,
		},
		// Negative numbers, with harmonious subsequence
		{
			name:     "Negative numbers",
			nums:     []int{-2, -3, -2, -1},
			expected: 2, // [-2, -3]
		},
		// All elements form harmonious subsequence
		{
			name:     "All numbers difference 1",
			nums:     []int{10, 9, 10, 9, 9},
			expected: 5, // [10,9,10,9,9]
		},
		// No two elements with difference 1
		{
			name:     "Non harmonious large gap",
			nums:     []int{10, 100, 1000},
			expected: 0,
		},
		// Multiple possible pairs, take max length
		{
			name:     "Multiple harmonious pairs",
			nums:     []int{1, 2, 2, 1, 3, 3, 3, 2, 4, 5},
			expected: 4, // [1,2,2,1] or [2,3,3,2] or similar
		},
		// Large array, all elements same
		{
			name:     "Large same number array",
			nums:     make([]int, 20000), // [0,0,0,...]
			expected: 0,
		},
	}

	// Fill the "Large same number array" with zeroes
	for i := range tests[9].nums {
		tests[9].nums[i] = 0
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findLHS(tc.nums) // Assume the function is implemented elsewhere
			if got != tc.expected {
				t.Errorf("For %v: expected %d, got %d", tc.nums, tc.expected, got)
			}
		})
	}
}
