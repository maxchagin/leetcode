package permutations

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function: sort each permutation and the outer slice for comparison
func normalize(output [][]int) [][]int {
	for i := range output {
		sort.Ints(output[i])
	}
	sort.Slice(output, func(i, j int) bool {
		for k := range output[i] {
			if k >= len(output[j]) {
				return false
			}
			if output[i][k] != output[j][k] {
				return output[i][k] < output[j][k]
			}
		}
		return len(output[i]) < len(output[j])
	})
	return output
}

// Helper function: check equality of two slices of integer slices ignoring order
func equalPermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	na := normalize(a)
	nb := normalize(b)
	return reflect.DeepEqual(na, nb)
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Single element",
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "Two elements",
			input: []int{1, 2},
			expected: [][]int{
				{1, 2}, {2, 1},
			},
		},
		{
			name:  "Three elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:  "Includes zero",
			input: []int{0, 1},
			expected: [][]int{
				{0, 1}, {1, 0},
			},
		},
		{
			name:  "Negative and positive numbers",
			input: []int{-1, 1},
			expected: [][]int{
				{-1, 1}, {1, -1},
			},
		},
		{
			name:  "Three with negative",
			input: []int{-1, 0, 1},
			expected: [][]int{
				{-1, 0, 1}, {-1, 1, 0},
				{0, -1, 1}, {0, 1, -1},
				{1, -1, 0}, {1, 0, -1},
			},
		},
		{
			name:  "Four elements ascending",
			input: []int{1, 2, 3, 4},
			expected: [][]int{
				// we only validate with equalPermutations, so we can put partial expected
				// But here to be precise it's 24 permutations in expected
				// For brevity, we won't manually list all of them but in real test they should be included
				// Example:
				{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2},
				{1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3},
				{2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
				{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1},
				{3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2},
				{4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
			},
		},
		{
			name:  "Three consecutive negatives",
			input: []int{-3, -2, -1},
			expected: [][]int{
				{-3, -2, -1}, {-3, -1, -2},
				{-2, -3, -1}, {-2, -1, -3},
				{-1, -3, -2}, {-1, -2, -3},
			},
		},
		{
			name:  "Max size input 6 elements",
			input: []int{1, 2, 3, 4, 5, 6},
			// Expected contains 720 permutations, we won’t list them explicitly.
			// Instead, we’ll assert only the count.
			expected: nil,
		},
		{
			name:  "Mixed small negatives and positives",
			input: []int{-1, 2, -3},
			expected: [][]int{
				{-1, 2, -3}, {-1, -3, 2},
				{2, -1, -3}, {2, -3, -1},
				{-3, -1, 2}, {-3, 2, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permute(tt.input)

			if tt.expected == nil {
				// Case with too many permutations -> only validate count
				expectedCount := 1
				for i := 2; i <= len(tt.input); i++ {
					expectedCount *= i
				}
				if len(result) != expectedCount {
					t.Errorf("Expected %d permutations, got %d", expectedCount, len(result))
				}
			} else {
				if !equalPermutations(result, tt.expected) {
					t.Errorf("Permutations mismatch.\nExpected: %v\nGot: %v", tt.expected, result)
				}
			}
		})
	}
}
