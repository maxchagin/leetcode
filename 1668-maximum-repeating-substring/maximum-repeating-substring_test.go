package maximumrepeatingsubstring

import "testing"

// TestMaxRepeating is a test function for maxRepeating function.
func TestMaxRepeating(t *testing.T) {
	tests := []struct {
		sequence string
		word     string
		expected int
	}{
		// Test case 1: Word is repeated twice in the sequence
		{"ababc", "ab", 2},

		// Test case 2: Word is repeated once in the sequence
		{"ababc", "ba", 1},

		// Test case 3: Word is not in the sequence
		{"ababc", "ac", 0},

		// Test case 4: Word "a" is repeated five times in a row
		{"aaaaa", "a", 5},

		// Test case 5: Word "ab" is repeated twice at the start of sequence
		{"abababab", "abab", 2},

		// Test case 6: Word "xyz" is not in the sequence, expected 0
		{"abcdefghijklmnop", "xyz", 0},

		// Test case 7: Empty word, expected 0
		{"abcdefghijklmnop", "", 0},

		// Test case 8: Empty sequence and non-empty word, expected 0
		{"", "a", 0},

		// Test case 9: Both sequence and word are empty, expected 0
		{"", "", 0},

		// Test case 10: Word "abc" is repeated non-contiguously, but consecutively as substring
		{"abcabcabc", "abc", 3},

		// Test case 11: the most discussed case!
		// (aaaba) aaab (aaaba aaaba aaaba aaaba aaaba)
		{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba", 5},
	}

	for i, test := range tests {
		// Execute the function with current test case inputs
		result := maxRepeating(test.sequence, test.word)

		// Check if the result matches the expected value
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d when sequence is '%s' and word is '%s'", i+1, test.expected, result, test.sequence, test.word)
		}
	}
}
