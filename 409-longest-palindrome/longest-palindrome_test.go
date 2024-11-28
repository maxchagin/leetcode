package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// Test case 1: Basic case with mixed characters
	input1 := "abccccdd"
	expected1 := 7
	// Explanation: One longest palindrome that can be built is "dccaccd"
	assertLongestPalindrome(t, input1, expected1)

	// Test case 2: Single character
	input2 := "a"
	expected2 := 1
	// Explanation: The longest palindrome that can be built is "a"
	assertLongestPalindrome(t, input2, expected2)

	// Test case 3: Mixed case with no possible palindromes longer than 1
	input3 := "abc"
	expected3 := 1
	// Explanation: Only single characters "a", "b", or "c" can be used as palindrome
	assertLongestPalindrome(t, input3, expected3)

	// Test case 4: All identical characters
	input4 := "aaaa"
	expected4 := 4
	// Explanation: Entire string can be a palindrome
	assertLongestPalindrome(t, input4, expected4)

	// Test case 5: Longer palindrome possible with odd count of a central character
	input5 := "aabbcc"
	expected5 := 6
	// Explanation: Longest palindrome is "abcba" or similar
	assertLongestPalindrome(t, input5, expected5)

	// Test case 6: Empty string
	input6 := ""
	expected6 := 0
	// Explanation: No characters to form a palindrome
	assertLongestPalindrome(t, input6, expected6)

	// Test case 7: All distinct characters
	input7 := "abcdef"
	expected7 := 1
	// Explanation: Only single characters can be chosen as palindrome
	assertLongestPalindrome(t, input7, expected7)

	// Test case 8: String with repeating patterns
	input8 := "abcabcabc"
	expected8 := 7
	// Explanation: Longest palindrome like "abcbaabc" or similar can be formed
	assertLongestPalindrome(t, input8, expected8)

	// Test case 9: Two identical characters
	input9 := "bb"
	expected9 := 2
	// Explanation: Both characters can be used to form a palindrome "bb"
	assertLongestPalindrome(t, input9, expected9)

	// Test case 10: Case sensitive example
	input10 := "AaBbCc"
	expected10 := 1
	// Explanation: Case sensitivity means no matched pairs for a palindrome longer than 1
	assertLongestPalindrome(t, input10, expected10)
}

func assertLongestPalindrome(t *testing.T, input string, expected int) {
	result := longestPalindrome(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}
