package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	// Define the test cases in a struct
	tests := []struct {
		input    string // Input string to test
		expected bool   // Expected result for the test
	}{
		{"A man, a plan, a canal: Panama", true}, // Example 1
		{"race a car", false},                    // Example 2
		{" ", true},                              // Example 3
		{"", true},                               // Edge case: empty string
		{"0P", false},                            // Edge case: mixed alphanumeric characters
		{"madam", true},                          // Simple palindrome
		{"Madam", true},                          // Case insensitivity check
		{"Step on no pets", true},                // Palindrome with spaces
		{"Was it a car or a cat I saw", true},    // Longer palindrome with spaces
		{"No 'x' in Nixon", true},                // Palindrome with punctuation
		{"Marge, let's \"[went].\" I await {news} telegram.", true},
	}

	// Iterate over each test case
	for i, test := range tests {
		result := isPalindrome(test.input) // Call the function with the test input
		// Assert the result with the expected output
		if result != test.expected {
			t.Errorf("Test %d: For input '%s', expected %v, but got %v",
				i, test.input, test.expected, result)
		}
	}
}
