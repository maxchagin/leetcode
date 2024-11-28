package guessnumberhigherorlower

import "testing"

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		n        int
		pick     int
		expected int
	}{
		{10, 6, 6},
		{1, 1, 1},
		{2, 1, 1},
		{100, 75, 75},
		{50, 50, 50},
		{1000, 50, 50},
		{20000, 256, 256},
		{2126753390, 1702766719, 1702766719},
	}

	for _, test := range tests {
		pick = test.pick // Set the picked number
		result := guessNumber(test.n)
		if result != test.expected {
			t.Errorf("guessNumber(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}
