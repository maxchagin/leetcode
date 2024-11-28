package findtheindexofthefirstoccurrenceinastring

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string // The main string where we need to search for the substring
		needle   string // The substring to search for in the haystack
		expected int    // The expected index of the first occurrence of needle in haystack
	}{
		{"sadbutsad", "sad", 0},       // Example 1: 'sad' occurs at index 0 and 6, return 0
		{"leetcode", "leeto", -1},     // Example 2: 'leeto' does not occur in 'leetcode', return -1
		{"hello", "ll", 2},            // 'll' starts at index 2
		{"aaaaa", "bba", -1},          // 'bba' does not occur in 'aaaaa', return -1
		{"", "", 0},                   // Both strings are empty, should return 0
		{"a", "", 0},                  // Empty needle should match at index 0
		{"", "a", -1},                 // Non-empty needle in empty haystack should return -1
		{"findtheindex", "index", 7},  // 'index' starts at index 7
		{"findtheindex", "fin", 0},    // 'fin' starts at index 0
		{"overlaptestest", "test", 7}, // 'test' starts at index 7 despite repetition
		{"mississippi", "issip", 4},   // 'issip' starts at index 4
	}

	// Iterate over each test case
	for i, test := range tests {
		result := strStr(test.haystack, test.needle) // Call the function with the test case inputs

		// Verify if the result matches the expected output
		if result != test.expected {
			t.Errorf("Test %d: For haystack = '%s' and needle = '%s', expected %d, but got %d",
				i, test.haystack, test.needle, test.expected, result)
		}
	}
}
