package distringmatch

import (
	"testing"
)

func isValidPermutation(s string, perm []int) bool {
	n := len(s)
	// Check length equals n+1
	if len(perm) != n+1 {
		return false
	}
	// Check that perm is a permutation of 0,...,n
	seen := make([]bool, n+1)
	for _, v := range perm {
		if v < 0 || v > n || seen[v] {
			return false
		}
		seen[v] = true
	}
	// Verify relative order conditions
	for i, ch := range s {
		if ch == 'I' && perm[i] >= perm[i+1] {
			return false
		}
		if ch == 'D' && perm[i] <= perm[i+1] {
			return false
		}
	}
	return true
}

func TestDiStringMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		// Test case 1: Basic alternating pattern
		{name: "Basic pattern - IDID", s: "IDID"},
		// Test case 2: All increasing
		{name: "All increasing - III", s: "III"},
		// Test case 3: Mixed with ending D
		{name: "Mixed pattern - DDI", s: "DDI"},
		// Test case 4: Single character I
		{name: "Single char I", s: "I"},
		// Test case 5: Single character D
		{name: "Single char D", s: "D"},
		// Test case 6: Longer pattern with multiple changes
		{name: "Longer pattern - IDII", s: "IDII"},
		// Test case 7: Alternating starting with D - DIDI", s: "DIDI"},
		// Test case 8: Two increases then one decrease
		{name: "Pattern - IID", s: "IID"},
		// Test case 9: Complex mix
		{name: "Complex pattern - DIIID", s: "DIIID"},
		// Test case 10: All decreasing
		{name: "All decreasing - DDDDD", s: "DDDDD"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			perm := diStringMatch(test.s)
			if !isValidPermutation(test.s, perm) {
				t.Errorf("diStringMatch(%q) returned an invalid permutation: %v", test.s, perm)
			}
		})
	}
}
