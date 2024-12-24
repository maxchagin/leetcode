package lettercombinationsofaphonenumber

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		// Test case 1: Basic two-digit combination
		{
			name:     "Two digits",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		// Test case 2: Empty string
		{
			name:     "Empty string",
			digits:   "",
			expected: []string{},
		},
		// Test case 3: Single digit
		{
			name:     "Single digit",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		// Test case 4: Four digits
		{
			name:   "Four digits",
			digits: "2345",
			expected: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil",
				"aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil",
				"afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil",
				"bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil",
				"begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil",
				"bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil",
				"cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil",
				"cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil",
				"cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil",
			},
		},
		// Test case 5: Digits with 7 and 9 (4 letters each)
		{
			name:     "Digits with 7",
			digits:   "27",
			expected: []string{"ap", "aq", "ar", "as", "bp", "bq", "br", "bs", "cp", "cq", "cr", "cs"},
		},
		// Test case 6: Same digit repeated
		{
			name:     "Same digit repeated",
			digits:   "22",
			expected: []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"},
		},
		// Test case 7: Three digits
		{
			name:     "Three digits",
			digits:   "234",
			expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
		// Test case 8: Single digit 9
		{
			name:     "Single digit 9",
			digits:   "9",
			expected: []string{"w", "x", "y", "z"},
		},
		// Test case 9: Digits 9 and 8
		{
			name:     "Digits 9 and 8",
			digits:   "98",
			expected: []string{"wt", "wu", "wv", "xt", "xu", "xv", "yt", "yu", "yv", "zt", "zu", "zv"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := letterCombinations(test.digits)
			// Sort both slices to compare them regardless of order
			sort.Strings(result)
			sort.Strings(test.expected)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("letterCombinations(%q) = %v, want %v",
					test.digits, result, test.expected)
			}
		})
	}
}
