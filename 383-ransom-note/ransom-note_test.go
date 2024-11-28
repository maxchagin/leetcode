package ransomnote

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"aa", "aabz", true},
		{"aa", "ab", false},
	}

	for i, test := range tests {
		result := canConstruct(test.ransomNote, test.magazine)
		if result != test.expected {
			t.Errorf("Test case %d failed: canConstruct(%v, %v) = %v; expected %v",
				i+1, test.ransomNote, test.magazine, result, test.expected)
		}
	}
}
