package issubsequence

import (
	"fmt"
	"testing"
)

// Function to test isSubsequence implementation
func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		// Test case 1: Simple subsequence
		{"abc", "ahbgdc", true},

		// Test case 2: Not a subsequence
		{"axc", "ahbgdc", false},

		// Test case 3: Subsequence with all characters present non-consecutively
		{"ace", "abcde", true},

		// Test case 4: Empty s string
		{"", "ahbgdc", true}, // An empty string is a subsequence of any string

		// Test case 5: Both s and t are empty
		{"", "", true}, // Both are empty, hence subsequence by default

		// Test case 6: s is longer than t
		{"abc", "ab", false}, // s cannot be a subsequence of a shorter t

		// Test case 7: s and t are the same
		{"hello", "hello", true}, // A string is always a subsequence of itself

		// Test case 8: Single character in s
		{"a", "a", true}, // Single identical character

		// Test case 9: Single character in s not in t
		{"z", "abcd", false}, // Character 'z' is not in t

		// Test case 10: s is a subsequence at the end of t
		{"fg", "abcdefg", true}, // 'fg' is a subsequence appearing at the end of t

		// Test case 11
		{"b", "abc", true},
	}

	for i, test := range tests {
		result := isSubsequence(test.s, test.t)
		if result != test.expected {
			fmt.Printf("Test case %d failed: got %v, expected %v\n", i+1, result, test.expected)
		} else {
			fmt.Printf("Test case %d passed\n", i+1)
		}
	}
}
