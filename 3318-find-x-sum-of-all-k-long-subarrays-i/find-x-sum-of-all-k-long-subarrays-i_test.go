package findxsumofallklongsubarraysi

import (
	"reflect"
	"testing"
)

func TestXSumOfKLongSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1 from problem",
			nums:     []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:        6,
			x:        2,
			expected: []int{6, 10, 12},
		},
		{
			name:     "Example 2 from problem",
			nums:     []int{3, 8, 7, 8, 7, 5},
			k:        2,
			x:        2,
			expected: []int{11, 15, 15, 15, 12},
		},
		{
			name:     "k equals x equals 1",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			x:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "k equals x equals array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			x:        5,
			expected: []int{15},
		},
		{
			name:     "x equals 1, keep most frequent element only",
			nums:     []int{1, 2, 2, 3, 3, 3, 4},
			k:        4,
			x:        1,
			expected: []int{4, 6, 9, 9},
		},
		{
			name:     "all elements same",
			nums:     []int{5, 5, 5, 5, 5},
			k:        3,
			x:        2,
			expected: []int{15, 15, 15},
		},
		{
			name:     "x greater than distinct elements",
			nums:     []int{1, 2, 1, 2, 3},
			k:        4,
			x:        3,
			expected: []int{6, 8},
		},
		{
			name:     "tie breaking by larger value",
			nums:     []int{1, 1, 2, 2, 3, 4},
			k:        4,
			x:        2,
			expected: []int{6, 7, 8},
		},
		{
			name:     "minimum constraints",
			nums:     []int{1},
			k:        1,
			x:        1,
			expected: []int{1},
		},
		{
			name:     "multiple ties in frequency",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
			k:        5,
			x:        2,
			expected: []int{6, 10, 10, 14, 14},
		},
		{
			name:     "x equals k but less than distinct elements",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        3,
			x:        3,
			expected: []int{6, 9, 12, 15},
		},
		{
			name:     "complex frequency distribution",
			nums:     []int{1, 2, 1, 3, 2, 1, 4, 3, 2, 1},
			k:        5,
			x:        2,
			expected: []int{6, 6, 6, 10, 8, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is a placeholder for the actual implementation
			result := findXSum(tt.nums, tt.k, tt.x)

			// Check if result has correct length
			expectedLength := len(tt.nums) - tt.k + 1
			if len(result) != expectedLength {
				t.Errorf("Result length mismatch: expected %d, got %d", expectedLength, len(result))
				return
			}

			// Compare expected and actual results using reflect.DeepEqual
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Result mismatch:\nExpected: %v\nGot:      %v", tt.expected, result)
			}
		})
	}
}
