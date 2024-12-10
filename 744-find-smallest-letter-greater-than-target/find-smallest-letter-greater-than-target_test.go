package findsmallestlettergreaterthantarget

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	// Test case 1: Target is less than all elements in letters
	letters1 := []byte{'c', 'f', 'j'}
	target1 := byte('a')
	expected1 := byte('c')
	if result := nextGreatestLetter(letters1, target1); result != expected1 {
		t.Errorf("Test case 1 failed: expected %c, got %c", expected1, result)
	}

	// Test case 2: Target is equal to the first letter
	letters2 := []byte{'c', 'f', 'j'}
	target2 := byte('c')
	expected2 := byte('f')
	if result := nextGreatestLetter(letters2, target2); result != expected2 {
		t.Errorf("Test case 2 failed: expected %c, got %c", expected2, result)
	}

	// Test case 3: Target is greater than all characters in letters
	letters3 := []byte{'x', 'x', 'y', 'y'}
	target3 := byte('z')
	expected3 := byte('x')
	if result := nextGreatestLetter(letters3, target3); result != expected3 {
		t.Errorf("Test case 3 failed: expected %c, got %c", expected3, result)
	}

	// Test case 4: All elements in letters are the same
	letters4 := []byte{'a', 'a', 'a'}
	target4 := byte('a')
	expected4 := byte('a')
	if result := nextGreatestLetter(letters4, target4); result != expected4 {
		t.Errorf("Test case 4 failed: expected %c, got %c", expected4, result)
	}

	// Test case 5: Wrap around to the start of the list
	letters5 := []byte{'c', 'd', 'e'}
	target5 := byte('e')
	expected5 := byte('c')
	if result := nextGreatestLetter(letters5, target5); result != expected5 {
		t.Errorf("Test case 5 failed: expected %c, got %c", expected5, result)
	}

	// Test case 6: Target is in between two letters
	letters6 := []byte{'a', 'b', 'd'}
	target6 := byte('b')
	expected6 := byte('d')
	if result := nextGreatestLetter(letters6, target6); result != expected6 {
		t.Errorf("Test case 6 failed: expected %c, got %c", expected6, result)
	}

	// Test case 7: Target occurs multiple times
	letters7 := []byte{'e', 'f', 'f', 'j'}
	target7 := byte('f')
	expected7 := byte('j')
	if result := nextGreatestLetter(letters7, target7); result != expected7 {
		t.Errorf("Test case 7 failed: expected %c, got %c", expected7, result)
	}

	// Test case 8: Target is greater than some but less than others
	letters8 := []byte{'m', 'n', 'o'}
	target8 := byte('n')
	expected8 := byte('o')
	if result := nextGreatestLetter(letters8, target8); result != expected8 {
		t.Errorf("Test case 8 failed: expected %c, got %c", expected8, result)
	}

	// Test case 9: Only two different characters and target between them
	letters9 := []byte{'g', 'k'}
	target9 := byte('h')
	expected9 := byte('k')
	if result := nextGreatestLetter(letters9, target9); result != expected9 {
		t.Errorf("Test case 9 failed: expected %c, got %c", expected9, result)
	}

	// Test case 10: Letters array contains letters at varied positions in the alphabet
	letters10 := []byte{'b', 'd', 'x', 'z'}
	target10 := byte('c')
	expected10 := byte('d')
	if result := nextGreatestLetter(letters10, target10); result != expected10 {
		t.Errorf("Test case 10 failed: expected %c, got %c", expected10, result)
	}
}
