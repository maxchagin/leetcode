package substringsofsizethreewithdistinctcharacters

import "testing"

func TestCountGoodSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Basic case with one good substring",
			s:    "xyzzaz",
			want: 1,
			// Explanation: Substrings are "xyz", "yzz", "zza", "zaz". Only "xyz" has all distinct characters.
		},
		{
			name: "Multiple good substrings",
			s:    "aababcabc",
			want: 4,
			// Explanation: Good substrings are "abc", "bca", "cab", and "abc".
		},
		{
			name: "No good substrings (all substrings have duplicates)",
			s:    "aaabbb",
			want: 0,
			// Explanation: Substrings are "aaa", "aab", "abb", "bbb". All have repeating characters.
		},
		{
			name: "Very short string (length = 3, all distinct)",
			s:    "abc",
			want: 1,
			// Explanation: The only substring "abc" has all distinct characters.
		},
		{
			name: "Very short string (length = 3, with duplicates)",
			s:    "aaa",
			want: 0,
			// Explanation: The only substring "aaa" has all repeating characters.
		},
		{
			name: "Long string with alternating good and bad substrings",
			s:    "abcabcabcabc",
			want: 10,
			// Explanation: Good substrings are all instances of "abc", "bca", "cab", etc.
		},
		{
			name: "String with length less than 3 (edge case)",
			s:    "ab",
			want: 0,
			// Explanation: No substrings of length 3 exist.
		},
		{
			name: "All characters are the same (long string)",
			s:    "zzzzzzzzzz",
			want: 0,
			// Explanation: All possible substrings of length 3 will be "zzz".
		},
		{
			name: "Mixed characters with some good substrings",
			s:    "abcddefghij",
			want: 7,
			// Explanation: Good substrings include "abc", "bcd", "def", "efg", "fgh", "ghi", "hij".
			// The substring "dde" is invalid due to repeating 'd'.
		},
		{
			name: "String with length exactly 100 (maximum constraint)",
			s:    repeatString("abcdefghij", 10), // Creates a 100-character string.
			want: 98,
			// Explanation: Every 3-letter substring in this string has distinct characters.
			// Given the pattern "abcdefghij" has no repeating characters, all 98 possible substrings are valid.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubstrings(tt.s); got != tt.want {
				t.Errorf("countGoodSubstrings(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// Helper function to repeat a string n times (used for the 100-length test case).
func repeatString(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
