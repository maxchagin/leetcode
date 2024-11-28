package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// Test case 1: Single character string
	// A single character is a palindrome.
	s := "a"
	expected := "a"
	if result := longestPalindrome(s); result != expected {
		t.Errorf("Test case 1 failed. Expected %s, got %s", expected, result)
	}

	// Test case 2: All characters are the same
	// The whole string is the longest palindrome.
	s = "aaaa"
	expected = "aaaa"
	if result := longestPalindrome(s); result != expected {
		t.Errorf("Test case 2 failed. Expected %s, got %s", expected, result)
	}

	// Test case 3: Simple palindrome in the middle
	// A straightforward palindrome in the string.
	s = "babad"
	expected1 := "bab" // Either "bab" or "aba" is a valid palindrome here.
	expected2 := "aba"
	result := longestPalindrome(s)
	if result != expected1 && result != expected2 {
		t.Errorf("Test case 3 failed. Expected %s or %s, got %s", expected1, expected2, result)
	}

	// // Test case 4: Even length palindrome
	// // Longest palindrome is of even length.
	// s = "cbbd"
	// expected = "bb"
	// if result := longestPalindrome(s); result != expected {
	// 	t.Errorf("Test case 4 failed. Expected %s, got %s", expected, result)
	// }

	// // Test case 5: No palindrome longer than 1
	// // Only individual characters are palindromes, return any single character.
	// s = "abcde"
	// expected = "a" // Any single character from the string could be a valid result.
	// result = longestPalindrome(s)
	// if len(result) != 1 || s[0] != result[0] {
	// 	t.Errorf("Test case 5 failed. Expected a single character from the string, got %s", result)
	// }

	// // Test case 6: Entire string is a palindrome
	// // The whole string should be returned.
	// s = "racecar"
	// expected = "racecar"
	// if result := longestPalindrome(s); result != expected {
	// 	t.Errorf("Test case 6 failed. Expected %s, got %s", expected, result)
	// }

	// // Test case 7: Palindrome at the start of the string
	// // Longest palindrome starts at the beginning.
	// s = "anana"
	// expected = "anana"
	// if result := longestPalindrome(s); result != expected {
	// 	t.Errorf("Test case 7 failed. Expected %s, got %s", expected, result)
	// }

	// // Test case 8: Palindrome at the end of the string
	// // Longest palindrome ends at the end.
	// s = "ananb"
	// expected = "ana"
	// if result := longestPalindrome(s); result != expected {
	// 	t.Errorf("Test case 8 failed. Expected %s, got %s", expected, result)
	// }

	// // Test case 9: String with mixed character classes
	// // Testing with special characters and numbers.
	// s = "a1bc22cb1a"
	// expected = "1bc22cb1"
	// if result := longestPalindrome(s); result != expected {
	// 	t.Errorf("Test case 9 failed. Expected %s, got %s", expected, result)
	// }

	// // Test case 10: Long string with multiple palindromes
	// // Select the longest palindrome from a lengthy string.
	// s = "abcbaabcd"
	// expected1 = "abcba"
	// expected2 = "baabc"
	// result = longestPalindrome(s)
	// if result != expected1 && result != expected2 {
	// 	t.Errorf("Test case 10 failed. Expected %s or %s, got %s", expected1, expected2, result)
	// }
}
