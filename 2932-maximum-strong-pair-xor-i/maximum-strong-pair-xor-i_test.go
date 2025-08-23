package maximumstrongpairxori

import "testing"

func TestMaximumStrongPairXOR(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Basic case with multiple strong pairs",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 7,
		},
		{
			name:     "Example 2: Two elements with same value pairs only",
			nums:     []int{10, 100},
			expected: 0,
		},
		{
			name:     "Example 3: Mixed values with specific strong pairs",
			nums:     []int{5, 6, 25, 30},
			expected: 7,
		},
		{
			name:     "Single element array",
			nums:     []int{42},
			expected: 0, // Only one element: pair with itself gives XOR 0
		},
		{
			name:     "All identical elements",
			nums:     []int{7, 7, 7, 7, 7},
			expected: 0, // All pairs are (7,7) with XOR 0
		},
		{
			name:     "Consecutive numbers with maximum XOR at end",
			nums:     []int{10, 11, 12, 13, 14, 15},
			expected: 7, // 15 XOR 8 (if 8 exists) or similar pattern
		},
		{
			name:     "Large numbers with valid strong pairs",
			nums:     []int{50, 75, 100, 125, 150},
			expected: 242, // Maximum possible XOR for numbers in this range
		},
		{
			name:     "Numbers with only one valid strong pair",
			nums:     []int{2, 4, 8, 16},
			expected: 24,
		},
		{
			name:     "Numbers where min constraint limits pairs",
			nums:     []int{1, 100, 101},
			expected: 1, // Only same-element pairs are valid: (1,1), (100,100), (101,101)
		},
		{
			name:     "Mixed positive numbers with various ranges",
			nums:     []int{3, 10, 27, 33, 42, 99},
			expected: 58, // Potential maximum XOR from valid strong pairs
		},
		{
			name:     "Edge case with minimum constraint values",
			nums:     []int{1, 1, 2, 3},
			expected: 3, // 1 XOR 2 = 3, and |1-2| = 1 <= min(1,2)=1
		},
		{
			name:     "Numbers forming multiple valid strong pairs with high XOR",
			nums:     []int{8, 16, 24, 32, 40},
			expected: 56, // Maximum XOR from valid pairs in this sequence
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumStrongPairXor(tt.nums)
			if result != tt.expected {
				t.Errorf("maximumStrongPairXOR(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}
