package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	// Define test cases structure
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1: Target found multiple times",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "Example 2: Target not found",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "Example 3: Empty array",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "Single element array - target found",
			nums:   []int{5},
			target: 5,
			want:   []int{0, 0},
		},
		{
			name:   "Single element array - target not found",
			nums:   []int{5},
			target: 8,
			want:   []int{-1, -1},
		},
		{
			name:   "Target at beginning of array",
			nums:   []int{2, 2, 2, 3, 4, 5},
			target: 2,
			want:   []int{0, 2},
		},
		{
			name:   "Target at end of array",
			nums:   []int{1, 2, 3, 4, 5, 5},
			target: 5,
			want:   []int{4, 5},
		},
		{
			name:   "All elements are target",
			nums:   []int{7, 7, 7, 7, 7},
			target: 7,
			want:   []int{0, 4},
		},
		{
			name:   "Target appears once in middle",
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 4,
			want:   []int{3, 3},
		},
		{
			name:   "Large numbers with negative values",
			nums:   []int{-100, -100, 0, 100, 100, 100},
			target: 100,
			want:   []int{3, 5},
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
