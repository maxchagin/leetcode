package smallestsubsequenceofdistinctcharacters

import "testing"

func TestSmallestSubsequence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		name     string
	}{
		{
			input:    "bcabc",
			expected: "abc",
			name:     "Example with multiple duplicates",
		},
		{
			input:    "cbacdcbc",
			expected: "acdb",
			name:     "Example with overlapping duplicates",
		},
		{
			input:    "abcd",
			expected: "abcd",
			name:     "String with all unique characters",
		},
		{
			input:    "aa",
			expected: "a",
			name:     "String with repeated single character",
		},
		{
			input:    "aabbcc",
			expected: "abc",
			name:     "String with multiple repeated characters",
		},
		{
			input:    "azbxcydwevf",
			expected: "abcdefvwxyz",
			name:     "String with non-contiguous unique characters",
		},
		{
			input:    "abacb",
			expected: "abc",
			name:     "Simple mix of duplicates and uniques",
		},
		{
			input:    "bbcaac",
			expected: "bac",
			name:     "String with complexities in order",
		},
		{
			input:    "",
			expected: "",
			name:     "Empty string input",
		},
		{
			input:    "zxyxxyzxyz",
			expected: "xyz",
			name:     "String with all same characters repeated in different order",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestSubsequence(tt.input)
			if result != tt.expected {
				t.Errorf("For input %q, expected %q but got %q", tt.input, tt.expected, result)
			}
		})
	}
}
