package poweroftwo

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// Test case 1: Basic power of 2
		{name: "Simple power of 2", input: 16, expected: true},
		// Test case 2: Smallest power of 2
		{name: "Number 1", input: 1, expected: true},
		// Test case 3: Large power of 2
		{name: "Large power of 2", input: 1024, expected: true},
		// Test case 4: Not a power of 2
		{name: "Not a power of 2", input: 3, expected: false},
		// Test case 5: Zero
		{name: "Zero", input: 0, expected: false},
		// Test case 6: Negative power of 2
		{name: "Negative power of 2", input: -16, expected: false},
		// Test case 8: Close to power of 2
		{name: "Close to power of 2", input: 33, expected: false},
		// Test case 9: Negative non-power of 2
		{name: "Negative non-power of 2", input: -7, expected: false},
		// Test case 10: Large non-power of 2
		{name: "Large non-power of 2", input: 1023, expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := isPowerOfTwo(test.input); result != test.expected {
				t.Errorf("isPowerOfTwo(%d) = %v, want %v",
					test.input, result, test.expected)
			}
		})
	}
}
