package findthekbeautyofanumber

import "testing"

func TestKBeauty(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			num:      240,
			k:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			num:      430043,
			k:        2,
			expected: 2,
		},
		{
			name:     "Single digit k=1",
			num:      7,
			k:        1,
			expected: 1, // "7" divides 7
		},
		{
			name:     "No k-beauty found",
			num:      123,
			k:        2,
			expected: 0, // "12" and "23" don't divide 123
		},
		{
			name:     "k equals number length",
			num:      1200,
			k:        4,
			expected: 1, // only "1200" itself divides 1200
		},
		{
			name:     "Number with leading zero substring",
			num:      1001,
			k:        2,
			expected: 1, // "10" divides 1001 (1001/10=100.1 but 1001%10=1), "00" invalid, "01" invalid
		},
		{
			name:     "Large number with multiple k-beauty",
			num:      123456789,
			k:        3,
			expected: 0,
		},
		{
			name:     "All zero substring",
			num:      10000,
			k:        3,
			expected: 1,
		},
		{
			name:     "k=1 with prime number",
			num:      13,
			k:        1,
			expected: 1,
		},
		{
			name:     "Large k-beauty with repeating pattern",
			num:      42424242,
			k:        2,
			expected: 4, // "42" divides 42424242 multiple times
		},
		{
			name:     "Edge case minimum values",
			num:      1,
			k:        1,
			expected: 1, // "1" divides 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorSubstrings(tt.num, tt.k); got != tt.expected {
				t.Errorf("divisorSubstrings() = %v, want %v", got, tt.expected)
			}
		})
	}
}
