package minimumdifferencebetweenhighestandlowestofkscores

import "testing"

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Single element array",
			nums:     []int{90},
			k:        1,
			expected: 0,
		},
		{
			name:     "Example from problem description",
			nums:     []int{9, 4, 1, 7},
			k:        2,
			expected: 2,
		},
		{
			name:     "All identical elements",
			nums:     []int{5, 5, 5, 5, 5},
			k:        3,
			expected: 0,
		},
		{
			name:     "Sorted array with k equal to array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: 4,
		},
		{
			name:     "Large range of numbers",
			nums:     []int{1, 100, 200, 300, 400, 500},
			k:        2,
			expected: 99,
		},
		{
			name:     "Unsorted array with minimum values",
			nums:     []int{0, 2, 5, 10, 15, 20},
			k:        3,
			expected: 5,
		},
		{
			name:     "Minimum k value",
			nums:     []int{10, 20, 30, 40},
			k:        1,
			expected: 0,
		},
		{
			name:     "Array with duplicate values",
			nums:     []int{3, 7, 3, 10, 7, 15, 7, 3},
			k:        4,
			expected: 3,
		},
		{
			name:     "Edge case with maximum constraint values",
			nums:     []int{0, 100000, 50000, 75000, 25000},
			k:        3,
			expected: 50000,
		},
		{
			name:     "Large unsorted array",
			nums:     []int{23, 45, 12, 67, 89, 34, 56, 78, 90, 11},
			k:        4,
			expected: 23,
		},
		{
			name:     "Leetcode 1",
			nums:     []int{87063, 61094, 44530, 21297, 95857, 93551, 9918},
			k:        6,
			expected: 74560,
		},
		{
			// name:     "Leetcode 2",
			nums:     []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526, 79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889, 65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785, 47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193, 35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284, 44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102, 18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478},
			k:        5,
			expected: 1428,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDifference(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("minimumDifference(%v, %d) = %d, expected %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
