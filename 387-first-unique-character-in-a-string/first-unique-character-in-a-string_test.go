package firstuniquecharacterinastring

import "testing"

// TestFirstUniqChar tests the firstUniqChar function with various cases.
func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// Basic case with the first character unique
		{"leetcode", 0},

		// First unique character in the middle
		{"loveleetcode", 2},

		// No unique characters present
		{"aabb", -1},

		// Single character string
		{"z", 0},

		// All characters are the same
		{"aaaa", -1},

		// Unique character at the end
		{"abaccabadz", 8},

		// Long string with unique character in middle
		{"aaaabbbdccceeef", 7},
	}

	for _, test := range tests {
		result := firstUniqChar(test.input)
		if result != test.expected {
			t.Errorf("firstUniqChar(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
