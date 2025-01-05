package findminimuminrotatedsortedarray

import "testing"

// TestFindMin tests the findMin function with various test cases
func TestFindMin(t *testing.T) {
	// Define test cases structure
	tests := []struct {
		name     string // Test case name
		input    []int  // Input array
		expected int    // Expected minimum value
	}{
		{
			name:     "LeetCode Example 1 - Array rotated 3 times",
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "LeetCode Example 2 - Array rotated 4 times",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "LeetCode Example 3 - Already sorted array",
			input:    []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "Single element array",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Two elements, rotated once",
			input:    []int{2, 1},
			expected: 1,
		},
		{
			name:     "Array with negative numbers",
			input:    []int{2, 3, 4, -5, -4, -3},
			expected: -5,
		},
		{
			name:     "Array rotated n-1 times",
			input:    []int{2, 1, 2, 2, 2},
			expected: 1,
		},
		{
			name:     "Array with maximum rotation",
			input:    []int{5000, -5000, -4999, -4998},
			expected: -5000,
		},
		{
			name:     "Array rotated once",
			input:    []int{2, 3, 4, 5, 6, 1},
			expected: 1,
		},
		{
			name:     "Array with all negative numbers",
			input:    []int{-3, -2, -1, -4},
			expected: -4,
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.input); got != tt.expected {
				t.Errorf("findMin() = %v, want %v", got, tt.expected)
			}
		})
	}
}
