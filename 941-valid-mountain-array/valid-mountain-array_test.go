package validmountainarray

import "testing"

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		// Test case 1: Fewer than 3 elements, not a mountain array
		{arr: []int{2, 1}, want: false},
		// Test case 2: Plateau at the peak, not strictly increasing then decreasing
		{arr: []int{3, 5, 5}, want: false},
		// Test case 3: Perfect mountain array
		{arr: []int{0, 3, 2, 1}, want: true},
		// Test case 4: Plateau at the start, not strictly increasing
		{arr: []int{5, 5, 3, 2, 1}, want: false},
		// Test case 5: Continuous increase without decrease
		{arr: []int{1, 2, 3, 4, 5}, want: false},
		// Test case 6: Continuous decrease without increase
		{arr: []int{5, 4, 3, 2, 1}, want: false},
		// Test case 7: Single peak at the end, requires at least three elements on each side
		{arr: []int{1, 2, 3}, want: false},
		// Test case 8: Single peak at the start, incorrect structure
		{arr: []int{3, 2, 0}, want: false},
		// Test case 9: Valid mountain array with distinct values
		{arr: []int{1, 3, 5, 4, 2, 0}, want: true},
		// Test case 10: Incorrect mountain with double peaks
		{arr: []int{0, 2, 3, 2, 3, 0}, want: false},
		{arr: []int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11}, want: true},
		{arr: []int{1, 3, 2}, want: true},
		{arr: []int{4, 20, 32, 45, 49, 45, 31, 21, 20, 16, 11, 8}, want: true},
	}

	for _, tt := range tests {
		got := validMountainArray(tt.arr)
		if got != tt.want {
			t.Errorf("validMountainArray(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}
