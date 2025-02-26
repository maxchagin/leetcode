package constructsmallestnumberfromdistring

import "testing"

func TestConstructSmallestNumberFromDIString(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected string
	}{
		// Test case 1: Example from description.
		{name: "Example 1", pattern: "IIIDIDDD", expected: "123549876"},
		// Test case 2: Example with all decreases.
		{name: "All Decreasing", pattern: "DDD", expected: "4321"},
		// Test case 3: Single increase.
		{name: "Single Increase", pattern: "I", expected: "12"},
		// Test case 4: Single decrease.
		{name: "Single Decrease", pattern: "D", expected: "21"},
		// Test case 5: Mixed pattern "ID".
		{name: "Mixed ID", pattern: "ID", expected: "132"},
		// Test case 6: Mixed pattern "DI".
		{name: "Mixed DI", pattern: "DI", expected: "213"},
		// Test case 7: Two consecutive increases.
		{name: "Two Increases", pattern: "II", expected: "123"},
		// Test case 8: Two consecutive decreases.
		{name: "Two Decreases", pattern: "DD", expected: "321"},
		// Test case 9: Alternating pattern "IDID".
		{name: "Alternating IDID", pattern: "IDID", expected: "13254"},
		// Test case 10: Alternating pattern "DIDI".
		{name: "Alternating DIDI", pattern: "DIDI", expected: "21435"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := constructSmallestNumberFromDIString(test.pattern)
			if result != test.expected {
				t.Errorf("constructSmallestNumberFromDIString(%q) = %q, want %q",
					test.pattern, result, test.expected)
			}
		})
	}
}
