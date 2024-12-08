package reversestring

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input  []byte
		output []byte
	}{
		// Test with a basic word containing all lowercase letters
		{input: []byte{'h', 'e', 'l', 'l', 'o'}, output: []byte{'o', 'l', 'l', 'e', 'h'}},

		// Test with a palindrome, a string that reads the same backward
		{input: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, output: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},

		// Test with a single character string
		{input: []byte{'a'}, output: []byte{'a'}},

		// Test with an empty string
		{input: []byte{}, output: []byte{}},

		// Test with a string containing both uppercase and lowercase letters
		{input: []byte{'A', 'b', 'C'}, output: []byte{'C', 'b', 'A'}},

		// Test with a string containing special characters and spaces
		{input: []byte{'$', ' ', '&'}, output: []byte{'&', ' ', '$'}},

		// Test with numbers as characters
		{input: []byte{'3', '2', '1'}, output: []byte{'1', '2', '3'}},

		// Test with mix of letters and numbers
		{input: []byte{'a', '1', 'b', '2'}, output: []byte{'2', 'b', '1', 'a'}},

		// Test with a longer sentence containing spaces
		{input: []byte{'I', ' ', 'a', 'm', ' ', 'h', 'e', 'r', 'e'}, output: []byte{'e', 'r', 'e', 'h', ' ', 'm', 'a', ' ', 'I'}},

		// Test with repeating characters
		{input: []byte{'x', 'x', 'x', 'x'}, output: []byte{'x', 'x', 'x', 'x'}},
	}

	for _, test := range tests {
		inputCopy := make([]byte, len(test.input))
		copy(inputCopy, test.input) // Make a copy to ensure the original test input is unchanged
		reverseString(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, inputCopy)
		}
	}
}
