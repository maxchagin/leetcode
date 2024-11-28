package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([{}])", true},
		{"(([]){})", true},
		{"(([]){})[", false},
		{"", true},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isValid(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
