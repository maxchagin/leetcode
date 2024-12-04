package makestringasubsequenceusingcyclicincrements

import "testing"

// TestCanMakeSubsequence tests the canMakeSubsequence function.
func TestCanMakeSubsequence(t *testing.T) {
	// Test case 1: Simple example where str2 is directly a subsequence after one operation.
	str1 := "abc"
	str2 := "ad"
	expected := true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected, result)
	}

	// Test case 2: str2 becomes a subsequence by transforming multiple characters.
	str1 = "zc"
	str2 = "ad"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected, result)
	}

	// Test case 3: Impossible to make str2 a subsequence of str1.
	str1 = "ab"
	str2 = "d"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected, result)
	}

	// Test case 4: When str1 is already str2 without any changes.
	str1 = "hello"
	str2 = "hello"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected, result)
	}

	// Test case 5: str2 is empty, should always return true.
	str1 = "anything"
	str2 = ""
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 5 failed: expected %v, got %v", expected, result)
	}

	// Test case 6: Large str1 with single character str2 that can be made by one operation.
	str1 = "zzzzzzzzzz"
	str2 = "a"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 6 failed: expected %v, got %v", expected, result)
	}

	// Test case 7: Identical single characters.
	str1 = "a"
	str2 = "a"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 7 failed: expected %v, got %v", expected, result)
	}

	// Test case 8: Impossible match due to character sequence.
	str1 = "abcde"
	str2 = "fgh"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 8 failed: expected %v, got %v", expected, result)
	}

	// Test case 9: Very large strings where transformation is not necessary.
	str1 = "abcd" + "efg" + "hijk" + "lmnop" // length: 100,000 is assumed
	str2 = "aceg"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 9 failed: expected %v, got %v", expected, result)
	}

	// Test case 10: Edge case with no possible transformations.
	str1 = "zzzzzzzz"
	str2 = "yyyyyyyy"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 10 failed: expected %v, got %v", expected, result)
	}

	// Test case 11
	str1 = "ww"
	str2 = "x"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 11 failed: expected %v, got %v", expected, result)
	}

	// Test case 12
	str1 = "eao"
	str2 = "ofa"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 12 failed: expected %v, got %v", expected, result)
	}

	// Test case 13
	str1 = "ff"
	str2 = "fg"
	expected = true
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 13 failed: expected %v, got %v", expected, result)
	}

	// Test case 14
	str1 = "abc"
	str2 = "abcd"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 14 failed: expected %v, got %v", expected, result)
	}

	// Test case 15
	str1 = "om"
	str2 = "nm"
	expected = false
	if result := canMakeSubsequence(str1, str2); result != expected {
		t.Errorf("Test case 15 failed: expected %v, got %v", expected, result)
	}
}
