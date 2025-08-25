package arithmeticslices

import "testing"

func TestArithmeticSlices(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Basic arithmetic sequence",
			nums:     []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "Example 2: Single element array",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "Three elements arithmetic sequence",
			nums:     []int{1, 3, 5},
			expected: 1,
		},
		{
			name:     "Three elements non-arithmetic sequence",
			nums:     []int{1, 2, 4},
			expected: 0,
		},
		{
			name:     "Long arithmetic sequence with constant difference",
			nums:     []int{7, 7, 7, 7, 7},
			expected: 6, // 3-length: 3, 4-length: 2, 5-length: 1
		},
		{
			name:     "Negative numbers arithmetic sequence",
			nums:     []int{3, -1, -5, -9},
			expected: 3, // [3,-1,-5], [-1,-5,-9], [3,-1,-5,-9]
		},
		{
			name:     "Mixed sequence with multiple arithmetic subarrays",
			nums:     []int{1, 2, 3, 4, 1, 2, 3},
			expected: 4,
		},
		{
			name:     "No arithmetic slices at all",
			nums:     []int{1, 2, 4, 8, 16},
			expected: 0,
		},
		{
			name:     "Two separate arithmetic sequences",
			nums:     []int{1, 2, 3, 5, 7, 9},
			expected: 4, // [1,2,3], [5,7,9], and two 3-length sequences
		},
		{
			name:     "Large array with alternating pattern",
			nums:     []int{1, 3, 1, 3, 1, 3, 1, 3},
			expected: 0, // No constant difference between consecutive elements
		},
		{
			name:     "Leetcode 9",
			nums:     []int{1, 2, 3, 8, 9, 10},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := numberOfArithmeticSlices(test.nums)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d for input %v", test.expected, result, test.nums)
			}
		})
	}
}
