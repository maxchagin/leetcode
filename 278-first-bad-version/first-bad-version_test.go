package firstbadversion

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		n        int
		bad      int
		expected int
	}{
		// Test case 1: Single version, which is bad
		{n: 1, bad: 1, expected: 1},
		// Test case 2: Multiple versions, first bad version is 4
		{n: 5, bad: 4, expected: 4},
		// Test case 3: Multiple versions, first bad version is 1
		{n: 5, bad: 1, expected: 1},
		// Test case 4: Multiple versions, first bad version is 5
		{n: 5, bad: 5, expected: 5},
		// Test case 5: Large number of versions, first bad version is 500
		{n: 1000, bad: 500, expected: 500},
		// Test case 6: Large number of versions, first bad version is 1
		{n: 1000, bad: 1, expected: 1},
		// Test case 7: Large number of versions, first bad version is 1000
		{n: 1000, bad: 1000, expected: 1000},
		// Test case 8: Edge case, first bad version is 2
		{n: 2, bad: 2, expected: 2},
		// Test case 9: Edge case, first bad version is 1
		{n: 2, bad: 1, expected: 1},
		// Test case 10: Large number of versions, first bad version is 999
		{n: 1000, bad: 999, expected: 999},
	}

	for _, test := range tests {
		// Set the bad version for current test case
		badVersion = test.bad

		result := firstBadVersion(test.n)
		if result != test.expected {
			t.Errorf("firstBadVersion(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}
