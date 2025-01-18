package countandsay

import "testing"

func TestCountAndSay(t *testing.T) {
	// Test cases structure
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Base case n=1",
			input:    1,
			expected: "1",
		},
		{
			name:     "Second sequence n=2",
			input:    2,
			expected: "11",
		},
		{
			name:     "Third sequence n=3",
			input:    3,
			expected: "21",
		},
		{
			name:     "Fourth sequence n=4",
			input:    4,
			expected: "1211",
		},
		{
			name:     "Fifth sequence n=5",
			input:    5,
			expected: "111221",
		},
		{
			name:     "Sixth sequence n=6",
			input:    6,
			expected: "312211",
		},
		{
			name:     "Seventh sequence n=7",
			input:    7,
			expected: "13112221",
		},
		{
			name:     "Eighth sequence n=8",
			input:    8,
			expected: "1113213211",
		},
		{
			name:     "Edge case - minimum n=1",
			input:    1,
			expected: "1",
		},
		{
			name:     "Large input n=10",
			input:    10,
			expected: "13211311123113112211",
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			result := CountAndSay(tt.input)

			// Then
			if result != tt.expected {
				t.Errorf("CountAndSay(%d) = %s; expected %s",
					tt.input, result, tt.expected)
			}
		})
	}
}
