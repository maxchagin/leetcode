package poweroffour

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// Test case 1: Basic power of 4
		{name: "Basic power of 4", input: 16, expected: true},
		// Test case 2: Smallest power of 4
		{name: "Number 1", input: 1, expected: true},
		// Test case 3: Large power of 4
		{name: "Large power of 4", input: 1048576, expected: true}, // 4^10
		// Test case 4: Power of 2 but not 4
		{name: "Power of 2 not 4", input: 32, expected: false},
		// Test case 5: Zero
		{name: "Zero", input: 0, expected: false},
		// Test case 6: Negative power of 4
		{name: "Negative number", input: -16, expected: false},
		// Test case 7: Another power of 4
		{name: "Another power of 4", input: 256, expected: true}, // 4^4
		// Test case 8: Close to power of 4
		{name: "Close to power of 4", input: 15, expected: false},
		// Test case 9: Not a power of 4
		{name: "Random number", input: 100, expected: false},
		// Test case 10: Maximum valid power of 4 in int32
		{name: "Maximum valid power of 4", input: 1073741824, expected: true}, // 4^15
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := isPowerOfFour(test.input); result != test.expected {
				t.Errorf("isPowerOfFour(%d) = %v, want %v",
					test.input, result, test.expected)
			}
		})
	}
}
