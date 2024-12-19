package checkifnumberisasumofpowersofthree

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// Test case 1: Basic sum of powers (12 = 3¹ + 3²)
		{name: "Basic sum 12", input: 12, expected: true},

		// Test case 2: Another valid sum (91 = 3⁰ + 3² + 3⁴)
		{name: "Valid sum 91", input: 91, expected: true},

		// Test case 3: Invalid number
		{name: "Invalid number 21", input: 21, expected: false},

		// Test case 4: Single power of three
		{name: "Single power 27", input: 27, expected: true},

		// Test case 5: Zero (edge case)
		{name: "Zero case", input: 0, expected: false},

		// Test case 6: Large valid number (3⁰ + 3¹ + 3² + 3³)
		{name: "Large valid sum", input: 40, expected: true},

		// Test case 7: Negative number
		{name: "Negative number", input: -9, expected: false},

		// Test case 8: Number that looks like power but isn't valid sum
		{name: "Invalid looking sum", input: 45, expected: false},

		// Test case 9: Large invalid number
		{name: "Large invalid number", input: 100, expected: false},

		// Test case 10: Simple valid case (3⁰ = 1)
		{name: "Smallest valid case", input: 1, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := checkPowersOfThree(test.input); result != test.expected {
				t.Errorf("checkPowersOfThree(%d) = %v, want %v",
					test.input, result, test.expected)
			}
		})
	}
}
