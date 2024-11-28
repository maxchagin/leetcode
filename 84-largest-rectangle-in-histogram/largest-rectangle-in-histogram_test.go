package largestrectangleinhistogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights  []int
		expected int
		name     string
	}{
		{
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
			name:     "Example case with multiple bars",
		},
		{
			heights:  []int{2, 4},
			expected: 4,
			name:     "Two bars of different height",
		},
		// {
		// 	heights:  []int{},
		// 	expected: 0,
		// 	name:     "Empty input array",
		// },
		// {
		// 	heights:  []int{5},
		// 	expected: 5,
		// 	name:     "Single bar",
		// },
		{
			heights:  []int{1, 1, 1, 1, 1},
			expected: 5,
			name:     "Uniform height bars",
		},
		// {
		// 	heights:  []int{1, 3, 2, 1, 2},
		// 	expected: 5,
		// 	name:     "Peak at the middle",
		// },
		// {
		// 	heights:  []int{2, 3, 4, 5, 6},
		// 	expected: 12,
		// 	name:     "Increasing heights",
		// },
		// {
		// 	heights:  []int{6, 5, 4, 3, 2},
		// 	expected: 12,
		// 	name:     "Decreasing heights",
		// },
		// {
		// 	heights:  []int{2, 2, 2, 2},
		// 	expected: 8,
		// 	name:     "All bars of equal height",
		// },
		// {
		// 	heights:  []int{2, 1, 2},
		// 	expected: 3,
		// 	name:     "Lowest bar in the middle",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.heights)
			if result != tt.expected {
				t.Errorf("For heights %v, expected %d but got %d", tt.heights, tt.expected, result)
			}
		})
	}
}
