package powerofthree

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// Test case 1: Basic power of 3
		{name: "Basic power of 3", input: 27, expected: true},
		// Test case 2: Smallest power of 3
		{name: "Number 1", input: 1, expected: true},
		// Test case 3: Large power of 3
		{name: "Large power of 3", input: 243, expected: true},
		// Test case 4: Not a power of 3
		{name: "Not a power of 3", input: 45, expected: false},
		// Test case 5: Zero
		{name: "Zero", input: 0, expected: false},
		// Test case 6: Negative power of 3
		{name: "Negative number", input: -27, expected: false},
		// Test case 7: Large non-power of 3
		{name: "Large non-power of 3", input: 100, expected: false},
		// Test case 8: Close to power of 3
		{name: "Close to power of 3", input: 26, expected: false},
		// Test case 9: Another power of 3
		{name: "Another power of 3", input: 9, expected: true},
		// Test case 10: Maximum power of 3 in int32
		{name: "Maximum valid power of 3", input: 1162261467, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := isPowerOfThree(test.input); result != test.expected {
				t.Errorf("isPowerOfThree(%d) = %v, want %v",
					test.input, result, test.expected)
			}
		})
	}
}
