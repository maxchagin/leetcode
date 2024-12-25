package smallestmissingintegergreaterthansequentialprefixsum

import "testing"

func TestMissingInteger(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Test case 1: Example from description
		{
			name:     "Basic example with break in sequence",
			nums:     []int{1, 2, 3, 2, 5},
			expected: 6,
		},
		// Test case 2: Another example from description
		{
			name:     "Sequence starting not from 1",
			nums:     []int{3, 4, 5, 1, 12, 14, 13},
			expected: 15,
		},
		// Test case 3: Single element array
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 6,
		},
		// Test case 4: All sequential numbers
		{
			name:     "Completely sequential array",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		// Test case 5: No sequence except first element
		{
			name:     "No sequence after first element",
			nums:     []int{1, 5, 3, 2, 4},
			expected: 6,
		},
		// Test case 6: Two element sequence
		{
			name:     "Two element sequence",
			nums:     []int{1, 2, 5, 7, 6},
			expected: 3,
		},
		// Test case 7: Sequence in middle
		{
			name:     "Sequence not at start",
			nums:     []int{5, 1, 2, 3, 10},
			expected: 6,
		},
		// Test case 8: Maximum values
		{
			name:     "Maximum allowed values",
			nums:     []int{48, 49, 50, 1, 2, 3},
			expected: 147,
		},
		// Test case 9: Short sequence with duplicates
		{
			name:     "Sequence with duplicates after",
			nums:     []int{1, 2, 2, 2, 2},
			expected: 3,
		},
		// Test case 10: Multiple short sequences
		{
			name:     "Multiple short sequences",
			nums:     []int{1, 2, 4, 5, 7, 8},
			expected: 3,
		},
		// Test case 11: Complex array with repeated numbers
		{
			name:     "Complex array with repeated numbers",
			nums:     []int{13, 4, 2, 2, 3, 4, 1, 8, 3, 7, 7, 7, 1, 6, 3},
			expected: 14,
		},
		// Test case 12: Sequence starting with large number
		{
			name:     "Large starting number sequence",
			nums:     []int{38, 43, 44},
			expected: 39,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := missingInteger(test.nums); result != test.expected {
				t.Errorf("missingInteger(%v) = %d, want %d",
					test.nums, result, test.expected)
			}
		})
	}
}
