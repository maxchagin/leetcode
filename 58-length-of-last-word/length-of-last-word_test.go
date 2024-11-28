package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
		{"b ", 1},
		{"    ", 0}, // edge case без слов
		{"today is a nice day", 3},
	}

	for _, test := range tests {
		result := lengthOfLastWord(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected length of last word %d, but got %d",
				test.input, test.expected, result)
		}
	}
}
