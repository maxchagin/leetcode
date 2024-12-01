package groupanagrams

import (
	"sort"
	"testing"
)

// TestGroupAnagrams tests the groupAnagrams function
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		// Edge case: Single empty string
		{
			input:    []string{""},
			expected: [][]string{{""}},
		},
		// Edge case: Single character string
		{
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		// No anagrams possible
		{
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		// All anagrams of the same word
		{
			input:    []string{"aaa", "aaa", "aaa"},
			expected: [][]string{{"aaa", "aaa", "aaa"}},
		},
		// Anagrams with different length strings
		{
			input:    []string{"", "b", "bb", "bbb"},
			expected: [][]string{{""}, {"b"}, {"bb"}, {"bbb"}},
		},
		// Mixed length of anagram words
		{
			input:    []string{"a", "aa", "aaa", "aaaa"},
			expected: [][]string{{"a"}, {"aa"}, {"aaa"}, {"aaaa"}},
		},
		// Large number of identical strings
		{
			input:    []string{"word", "word", "word", "word"},
			expected: [][]string{{"word", "word", "word", "word"}},
		},
		{
			input:    []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			expected: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}},
		},
	}

	for _, test := range tests {
		output := groupAnagrams(test.input)
		if !isEqual(output, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, output)
		}
	}
}

// isEqual checks if two 2D slices contain the same elements, ignoring the order.
func isEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Normalize the order of elements in each group
	for i := range a {
		sort.Strings(a[i])
	}
	for i := range b {
		sort.Strings(b[i])
	}

	// Sort the groups themselves
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})

	// Compare normalized versions
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
