package nextgreaterelementiii

import "testing"

func TestNextGreaterElement(t *testing.T) {
	// Test case 1: Regular case with a small number.
	n := 12
	expected := 21
	result := nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 1 failed. Expected %d, got %d", expected, result)
	}

	// Test case 2: No greater permutation possible.
	n = 21
	expected = -1
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 2 failed. Expected %d, got %d", expected, result)
	}

	// Test case 3: Largest number edge case.
	n = 987654321
	expected = -1 // No greater number can be formed as it's already maximal
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 3 failed. Expected %d, got %d", expected, result)
	}

	// Test case 4: Medium-sized number with a valid greater permutation.
	n = 12443322
	expected = 13222344
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 4 failed. Expected %d, got %d", expected, result)
	}

	// Test case 5: Number with all identical digits.
	n = 111
	expected = -1 // All permutations are the same
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 5 failed. Expected %d, got %d", expected, result)
	}

	// Test case 6: Number with only two digits.
	n = 45
	expected = 54
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 6 failed. Expected %d, got %d", expected, result)
	}

	// Test case 7: Edge case where next permutation requires rearranging several digits.
	n = 29999999
	expected = -1
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 7 failed. Expected %d, got %d", expected, result)
	}

	// Test case 8: Number where the least significant digits form the largest permutation.
	n = 1999999999
	expected = -1 // This number is the max permutation and should return -1
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 8 failed. Expected %d, got %d", expected, result)
	}

	// Test case 9: Simple small number
	n = 9
	expected = -1 // Single digit, no greater number possible
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 9 failed. Expected %d, got %d", expected, result)
	}

	// Test case 10: Large valid permutation.
	n = 2017
	expected = 2071
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 10 failed. Expected %d, got %d", expected, result)
	}
	// Test case 11: Large valid permutation.
	n = 1234
	expected = 1243
	result = nextGreaterElement(n)
	if result != expected {
		t.Errorf("Test case 10 failed. Expected %d, got %d", expected, result)
	}
}
