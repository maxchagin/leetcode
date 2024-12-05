package movepiecestoobtainastring

import "testing"

func TestCanChange(t *testing.T) {
	cases := []struct {
		start   string
		target  string
		want    bool
		comment string
	}{
		// Test case 1: Example 1 from the prompt
		{"_L__R__R_", "L______RR", true, "Example 1: Possible to transform via multiple moves"},

		// Test case 2: Example 2 from the prompt
		{"R_L_", "__LR", false, "Example 2: 'R' cannot move left, 'L' cannot move right, transformation is impossible"},

		// Test case 3: Example 3 from the prompt
		{"_R", "R_", false, "Example 3: 'R' cannot move left due to lack of trailing space"},

		// Test case 4: No movement needed, start is equal to target
		{"LR_", "LR_", true, "Same strings: Already in the desired form"},

		// Test case 6: Shifting all 'R's right and all 'L's left
		{"_RR__LL_", "R___RL__", false, "Impossible since 'R' should move left or 'L' should move right"},

		// Test case 7: Large input size with impossible transformation
		{"R__R__L___", "__RR__L___", true, "All 'R' pieces cannot align due to initial leftmost position constraints"},

		// Test case 8: Single character difference making transformation impossible
		{"_L_", "_R_", false, "'L' cannot become 'R'"},

		// Test case 9: Switching positions without crossing
		{"RL_", "LR_", false, "'R' and 'L' cannot cross or switch positions"},
	}

	for i, c := range cases {
		got := canChange(c.start, c.target)
		if got != c.want {
			t.Errorf("Test case %d failed: %s, got %v, want %v", i+1, c.comment, got, c.want)
		}
	}
}
