package maxareaofisland

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1 from problem description",
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			expected: 6,
		},
		{
			name:     "Example 2 from problem description - all water",
			grid:     [][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
			expected: 0,
		},
		{
			name:     "Single cell island",
			grid:     [][]int{{1}},
			expected: 1,
		},
		{
			name:     "Single cell water",
			grid:     [][]int{{0}},
			expected: 0,
		},
		{
			name: "Multiple isolated islands of size 1",
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			expected: 1,
		},
		{
			name: "Large contiguous island",
			grid: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: 12,
		},
		{
			name: "Island with maximum possible area (50x50 grid of 1's)",
			grid: func() [][]int {
				grid := make([][]int, 50)
				for i := range grid {
					grid[i] = make([]int, 50)
					for j := range grid[i] {
						grid[i][j] = 1
					}
				}
				return grid
			}(),
			expected: 2500,
		},
		{
			name: "Diagonal ones that are not connected 4-directionally",
			grid: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 1,
		},
		{
			name: "Two islands of different sizes",
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Spiral shaped island",
			grid: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			},
			expected: 8,
		},
		{
			name: "Leetcode 609",
			grid: [][]int{
				{0, 1},
				{1, 1},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.grid); got != tt.expected {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.expected)
			}
		})
	}
}
