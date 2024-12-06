package maximumnumberofintegerstochoosefromarangei

import "testing"

func TestMaxCount(t *testing.T) {
	tests := []struct {
		banned []int
		n      int
		maxSum int
		want   int
	}{
		// Test case 1: Basic scenario where you can choose 2 and 4
		{banned: []int{1, 6, 5}, n: 5, maxSum: 6, want: 2},
		// Test case 2: All possible numbers are banned
		{banned: []int{1, 2, 3, 4, 5, 6, 7}, n: 8, maxSum: 1, want: 0},
		// Test case 3: No numbers are banned and maxSum is large enough
		{banned: []int{11}, n: 7, maxSum: 50, want: 7},
		// Test case 4: No numbers are banned and exact maxSum
		{banned: []int{}, n: 5, maxSum: 15, want: 5},
		// Test case 5: Large maxSum, but all small numbers are banned
		{banned: []int{1, 2, 3}, n: 5, maxSum: 100, want: 2},
		// Test case 6: Only one number available to choose
		{banned: []int{1, 2, 3, 4}, n: 5, maxSum: 1, want: 0},
		// Test case 7: Large n but all numbers just under maxSum
		{banned: []int{8, 9, 10}, n: 10, maxSum: 15, want: 5},
		// Test case 8: edge cases of minimum numbers
		{banned: []int{9999, 10000}, n: 10000, maxSum: 100, want: 13},
		// Test case 9: Only very large numbers are banned
		{banned: []int{10000}, n: 9999, maxSum: 5000, want: 99},
		// Test case 10: Only the maximum number in the range is banned
		{banned: []int{5}, n: 5, maxSum: 10, want: 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maxCount(tt.banned, tt.n, tt.maxSum)
			if got != tt.want {
				t.Errorf("maxCount(%v, %d, %d) = %d; want %d", tt.banned, tt.n, tt.maxSum, got, tt.want)
			}
		})
	}
}
