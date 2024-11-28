package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{9, 55},
		{10, 89},
	}

	for i, test := range tests {
		result := climbStairs(test.input)
		if result != test.expected {
			t.Errorf("Test %d: For input %d, expected %d, but got %d", i, test.input, test.expected, result)
		}
	}
}
