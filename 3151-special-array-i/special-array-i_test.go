package specialarrayi

import "testing"

func TestIsArraySpecial(t *testing.T) {
	cases := []struct {
		nums    []int
		want    bool
		comment string
	}{
		// Test case 1: Single element array
		{[]int{1}, true, "Single element, automatically special"},

		// Test case 2: Example 2 from the prompt
		{[]int{2, 1, 4}, true, "Pairs (2,1) and (1,4) have different parity"},

		// Test case 3: Example 3 from the prompt
		{[]int{4, 3, 1, 6}, false, "Pair (3,1) have the same parity, so not special"},

		// Test case 4: All even numbers
		{[]int{2, 4, 6, 8}, false, "All elements have the same parity"},

		// Test case 5: All odd numbers
		{[]int{1, 3, 5, 7}, false, "All elements have the same parity"},

		// Test case 6: Alternating even and odd
		{[]int{1, 2, 3, 4, 5}, true, "Even and odd elements alternate"},

		// Test case 7: Two elements of different parity
		{[]int{5, 6}, true, "One pair of different parity numbers"},

		// Test case 8: Large array with alternating parity
		{[]int{2, 3, 4, 5, 6, 7, 8, 9}, true, "Larger array with alternating even and odd numbers"},

		// Test case 9: Parities match at the end only
		{[]int{2, 1, 4, 3, 8, 6}, false, "Last two elements (7,6) are both odd"},
	}

	for i, c := range cases {
		got := isArraySpecial(c.nums)
		if got != c.want {
			t.Errorf("Test case %d failed: %s, got %v, want %v", i+1, c.comment, got, c.want)
		}
	}
}
