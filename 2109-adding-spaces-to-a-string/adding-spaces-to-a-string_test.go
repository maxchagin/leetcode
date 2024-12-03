package addingspacestoastring

import "testing"

// TestAddSpaces will test the addSpaces function.
func TestAddSpaces(t *testing.T) {
	// Test case 1: Standard case with multiple spaces to be added.
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}
	expected := "Leetcode Helps Me Learn"
	result := addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 1 failed: expected %q, got %q", expected, result)
	}

	// Test case 2: Different standard input.
	s = "icodeinpython"
	spaces = []int{1, 5, 7, 9}
	expected = "i code in py thon"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 2 failed: expected %q, got %q", expected, result)
	}

	// Test case 3: Spaces before each character including the first.
	s = "spacing"
	spaces = []int{0, 1, 2, 3, 4, 5, 6}
	expected = " s p a c i n g"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 3 failed: expected %q, got %q", expected, result)
	}

	// Test case 4: No spaces to be added.
	s = "nospaces"
	spaces = []int{}
	expected = "nospaces"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 4 failed: expected %q, got %q", expected, result)
	}

	// Test case 5: Spaces at the end of the string.
	// s = "endspacing"
	// spaces = []int{10}
	// expected = "endspacing "
	// result = addSpaces(s, spaces)
	// if result != expected {
	// 	t.Errorf("Test case 5 failed: expected %q, got %q", expected, result)
	// }

	// Test case 6: Only one space added, at the start.
	s = "single"
	spaces = []int{0}
	expected = " single"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 6 failed: expected %q, got %q", expected, result)
	}

	// Test case 7: Empty string.
	s = ""
	spaces = []int{}
	expected = ""
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 7 failed: expected %q, got %q", expected, result)
	}

	// Test case 8: Single character string with space.
	s = "x"
	spaces = []int{0}
	expected = " x"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 9 failed: expected %q, got %q", expected, result)
	}

	// Test case 9: Large string with random spaces.
	s = "largeexampleteststring"
	spaces = []int{5, 12, 16}
	expected = "large example test string"
	result = addSpaces(s, spaces)
	if result != expected {
		t.Errorf("Test case 10 failed: expected %q, got %q", expected, result)
	}
}
