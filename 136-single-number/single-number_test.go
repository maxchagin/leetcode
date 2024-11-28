package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	// Test case 1: Basic example
	nums1 := []int{2, 2, 1}
	expected1 := 1
	if result := singleNumber(nums1); result != expected1 {
		t.Errorf("Test case 1 failed. Expected %d, got %d", expected1, result)
	}

	// Test case 2: Another basic example
	nums2 := []int{4, 1, 2, 1, 2}
	expected2 := 4
	if result := singleNumber(nums2); result != expected2 {
		t.Errorf("Test case 2 failed. Expected %d, got %d", expected2, result)
	}

	// Test case 3: Single element
	nums3 := []int{1}
	expected3 := 1
	if result := singleNumber(nums3); result != expected3 {
		t.Errorf("Test case 3 failed. Expected %d, got %d", expected3, result)
	}

	// Test case 4: All negative numbers
	nums4 := []int{-1, -1, -2}
	expected4 := -2
	if result := singleNumber(nums4); result != expected4 {
		t.Errorf("Test case 4 failed. Expected %d, got %d", expected4, result)
	}

	// Test case 5: Mix of negative and positive numbers
	nums5 := []int{-1, 2, -1, 3, 2}
	expected5 := 3
	if result := singleNumber(nums5); result != expected5 {
		t.Errorf("Test case 5 failed. Expected %d, got %d", expected5, result)
	}

	// Test case 6: Large numbers
	nums6 := []int{100000, 100000, 999999}
	expected6 := 999999
	if result := singleNumber(nums6); result != expected6 {
		t.Errorf("Test case 6 failed. Expected %d, got %d", expected6, result)
	}

	// Test case 7: Zero as a single number
	nums7 := []int{0, 1, 1}
	expected7 := 0
	if result := singleNumber(nums7); result != expected7 {
		t.Errorf("Test case 7 failed. Expected %d, got %d", expected7, result)
	}

	// Test case 8: Single element repeated
	nums8 := []int{-3, -3, 9, 9, 15}
	expected8 := 15
	if result := singleNumber(nums8); result != expected8 {
		t.Errorf("Test case 8 failed. Expected %d, got %d", expected8, result)
	}

	// Test case 9: Multiple zeroes
	nums9 := []int{0, 0, 0, 0, 1}
	expected9 := 1
	if result := singleNumber(nums9); result != expected9 {
		t.Errorf("Test case 9 failed. Expected %d, got %d", expected9, result)
	}

	// Test case 10: Large array with single number at the end
	nums10 := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 7, 6, 6}
	expected10 := 7
	if result := singleNumber(nums10); result != expected10 {
		t.Errorf("Test case 10 failed. Expected %d, got %d", expected10, result)
	}
}
