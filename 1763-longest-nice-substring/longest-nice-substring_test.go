package longestnicesubstring

import (
	"testing"
)

func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Case 1: Example from the problem, longest nice substring in the middle
		{"YazaAay", "aAa"},
		// // Case 2: Whole string is nice
		// {"Bb", "Bb"},
		// // Case 3: No nice substring
		// {"c", ""},
		// // Case 4: Multiple nice substrings, return earliest
		// {"aAaBb", "aAa"},
		// // Case 5: All letters present in both cases
		// {"AaBbCc", "AaBbCc"},
		// // Case 6: Only one nice substring at the end
		// {"xYyZz", "YyZz"},
		// // Case 7: Nice substring at the beginning
		// {"AaXx", "Aa"},
		// // Case 8: Mixed cases, longest nice substring in the middle
		// {"abABcdCD", "abABcdCD"},
		// // Case 9: Multiple nice substrings of same length, return earliest
		// {"aA bB", "aA"},
		// // Case 10: String with no uppercase letters
		// {"abcdef", ""},
		// // Case 11: String with no lowercase letters
		// {"ABCDEF", ""},
		// // Case 12: String with alternating cases, all nice
		// {"aAbBcCdD", "aAbBcCdD"},
		// // Case 13: Nice substring with repeated letters
		// {"aAaAaA", "aAaAaA"},
		// // Case 14: Nice substring with only two letters
		// {"xX", "xX"},
		// // Case 15: Empty string
		// {"", ""},
	}

	for _, tc := range tests {
		result := longestNiceSubstring(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", tc.input, tc.expected, result)
		}
	}
}
