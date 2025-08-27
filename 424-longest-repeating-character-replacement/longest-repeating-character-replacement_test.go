package longestrepeatingcharacterreplacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Example 1: Basic replacement with k=2",
			s:        "ABAB",
			k:        2,
			expected: 4,
		},
		{
			name:     "Example 2: Single replacement with k=1",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "All same characters, no replacement needed",
			s:        "AAAAA",
			k:        2,
			expected: 5,
		},
		{
			name:     "Empty string",
			s:        "",
			k:        0,
			expected: 0,
		},
		{
			name:     "k=0, no replacements allowed",
			s:        "AABA",
			k:        0,
			expected: 2, // Longest consecutive same characters
		},
		{
			name:     "k equal to number of different characters",
			s:        "ABCABC",
			k:        3,
			expected: 5, // Can replace all to make same
		},
		{
			name:     "Large k value",
			s:        "ABCDEF",
			k:        10,
			expected: 6, // Can replace all characters to same
		},
		{
			name:     "Single character string",
			s:        "A",
			k:        0,
			expected: 1,
		},
		{
			name:     "All different characters with limited k",
			s:        "ABCDE",
			k:        2,
			expected: 3, // Can replace 2 characters to match others
		},
		{
			name:     "Mixed case with optimal replacement strategy",
			s:        "AABAABA",
			k:        1,
			expected: 5, // Should find optimal substring to replace
		},
		{
			name:     "Leetcode 40",
			s:        "ABBB",
			k:        2,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.expected {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
