package findtheprefixcommonarrayoftwoarrays

import (
	"reflect"
	"testing"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	tests := []struct {
		name     string
		A        []int
		B        []int
		expected []int
	}{
		// Test case 1: Example from description
		{
			name:     "Basic example from description",
			A:        []int{1, 3, 2, 4},
			B:        []int{3, 1, 2, 4},
			expected: []int{0, 2, 3, 4},
		},
		// Test case 2: Another example from description
		{
			name:     "Another basic example",
			A:        []int{2, 3, 1},
			B:        []int{3, 1, 2},
			expected: []int{0, 1, 3},
		},
		// Test case 3: Single element arrays
		{
			name:     "Single element arrays",
			A:        []int{1},
			B:        []int{1},
			expected: []int{1},
		},
		// Test case 4: Identical arrays
		{
			name:     "Identical arrays",
			A:        []int{1, 2, 3, 4},
			B:        []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		// Test case 5: Reverse order arrays
		{
			name:     "Reverse order arrays",
			A:        []int{1, 2, 3, 4},
			B:        []int{4, 3, 2, 1},
			expected: []int{0, 0, 2, 4},
		},
		// Test case 6: No common elements until end
		{
			name:     "No common elements until end",
			A:        []int{1, 2, 3},
			B:        []int{3, 1, 2},
			expected: []int{0, 1, 3},
		},
		// Test case 7: Larger arrays
		{
			name:     "Larger arrays",
			A:        []int{1, 2, 3, 4, 5},
			B:        []int{5, 4, 3, 2, 1},
			expected: []int{0, 0, 1, 3, 5},
		},
		// Test case 8: All elements different until last
		{
			name:     "All different until last",
			A:        []int{2, 3, 4, 1},
			B:        []int{4, 1, 3, 2},
			expected: []int{0, 0, 2, 4},
		},
		// Test case 9: Alternating pattern
		{
			name:     "Alternating pattern",
			A:        []int{1, 2, 3, 4},
			B:        []int{2, 1, 4, 3},
			expected: []int{0, 2, 2, 4},
		},
		// Test case 10: Maximum length case
		{
			name:     "Maximum length case",
			A:        []int{1, 2, 3, 4, 5, 6},
			B:        []int{6, 5, 4, 3, 2, 1},
			expected: []int{0, 0, 0, 2, 4, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findThePrefixCommonArray(test.A, test.B)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("findThePrefixCommonArray(%v, %v) = %v, want %v",
					test.A, test.B, result, test.expected)
			}
		})
	}
}
