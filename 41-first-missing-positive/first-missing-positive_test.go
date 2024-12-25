package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Test case 1: Example from description
		{
			name:     "Basic example with zero",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		// Test case 2: Array with negative number
		{
			name:     "Array with negative and missing 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		// Test case 3: All numbers are too large
		{
			name:     "All numbers greater than 1",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		// Test case 4: Single element array
		{
			name:     "Single element array",
			nums:     []int{2},
			expected: 1,
		},
		// Test case 5: Array with duplicates
		{
			name:     "Array with duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: 4,
		},
		// Test case 6: Consecutive positive integers
		{
			name:     "Consecutive positive integers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		// Test case 7: All negative numbers
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3},
			expected: 1,
		},
		// Test case 8: Mix of positive, negative and zero
		{
			name:     "Mixed numbers with gaps",
			nums:     []int{3, 4, -1, 1, 0, -5, 5, 4},
			expected: 2,
		},
		// Test case 9: Large numbers with small missing
		{
			name:     "Large numbers with small missing",
			nums:     []int{1000, 2000, 3000, 1, 3},
			expected: 2,
		},
		// Test case 10: Missing in the middle of sequence
		{
			name:     "Missing in sequence",
			nums:     []int{1, 2, 3, 5, 6},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := firstMissingPositive(test.nums); result != test.expected {
				t.Errorf("firstMissingPositive(%v) = %d, want %d",
					test.nums, result, test.expected)
			}
		})
	}
}
