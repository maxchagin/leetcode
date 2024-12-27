package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two houses",
			nums:     []int{2, 3},
			expected: 3,
		},
		{
			name:     "Three houses",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "Three houses 2",
			nums:     []int{1, 3, 1},
			expected: 3,
		},
		{
			name:     "14",
			nums:     []int{4, 1, 2, 7, 5, 3, 1},
			expected: 14,
		},
		{
			name:     "3",
			nums:     []int{1, 2, 1, 1},
			expected: 3,
		},
		{
			name:     "4",
			nums:     []int{2, 1, 1, 2},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.expected {
				t.Errorf("rob() = %v, want %v", got, tt.expected)
			}
		})
	}
}
