package countsubstringsthatsatisfykconstrainti

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Example 1: Basic case with k=1",
			s:        "10101",
			k:        1,
			expected: 12,
		},
		{
			name:     "Example 2: Longer string with k=2",
			s:        "1010101",
			k:        2,
			expected: 25,
		},
		{
			name:     "Example 3: All ones with k=1",
			s:        "11111",
			k:        1,
			expected: 15,
		},
		{
			name:     "Single character 0 with k=1",
			s:        "0",
			k:        1,
			expected: 1,
		},
		{
			name:     "Single character 1 with k=1",
			s:        "1",
			k:        1,
			expected: 1,
		},
		{
			name:     "All zeros with k=1",
			s:        "00000",
			k:        1,
			expected: 15,
		},
		{
			name:     "Alternating pattern with k=3",
			s:        "1010101010",
			k:        3,
			expected: 49,
		},
		{
			name:     "Large k value equal to string length",
			s:        "10101",
			k:        5,
			expected: 15,
		},
		{
			name:     "String with equal zeros and ones, k=2",
			s:        "110011",
			k:        2,
			expected: 21,
		},
		{
			name:     "Minimum length string with k=1",
			s:        "10",
			k:        1,
			expected: 3,
		},
		{
			name:     "String with many zeros and small k",
			s:        "000111000",
			k:        2,
			expected: 35,
		},
		{
			name:     "Edge case: maximum k for given string",
			s:        "111000",
			k:        6,
			expected: 21,
		},
		{
			name:     "Leetcode 1",
			s:        "011",
			k:        1,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings(tt.s, tt.k); got != tt.expected {
				t.Errorf("countSubstrings(%q, %d) = %d, Expected %d", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
