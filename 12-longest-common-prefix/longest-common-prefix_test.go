package longestcommonprefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"code", "code", "code"}, "code"},
		{[]string{"num123", "num456", "num789"}, "num"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
		{[]string{}, ""},
	}

	fmt.Println("-----------")

	for _, test := range tests {
		result := longestCommonPrefix(test.strs)
		if result != test.expected {
			t.Errorf("For input %v, expected %s but got %s", test.strs, test.expected, result)
		}
	}
}
