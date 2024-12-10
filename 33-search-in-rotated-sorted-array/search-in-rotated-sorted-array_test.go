package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	// Test case 1: Target is the pivot
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	expected1 := 4
	if result := search(nums1, target1); result != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result)
	}

	// Test case 2: Target is in the right part of the rotated array
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 1
	expected2 := 5
	if result := search(nums2, target2); result != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result)
	}

	// Test case 3: Target is not present in the array
	nums3 := []int{4, 5, 6, 7, 0, 1, 2}
	target3 := 3
	expected3 := -1
	if result := search(nums3, target3); result != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result)
	}

	// Test case 4: Target is the smallest element
	nums4 := []int{4, 5, 6, 7, 0, 1, 2}
	target4 := 0
	expected4 := 4
	if result := search(nums4, target4); result != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result)
	}

	// Test case 5: Target is the largest element
	nums5 := []int{4, 5, 6, 7, 0, 1, 2}
	target5 := 7
	expected5 := 3
	if result := search(nums5, target5); result != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result)
	}

	// Test case 6: The array is not rotated
	nums6 := []int{1, 2, 3, 4, 5, 6, 7}
	target6 := 3
	expected6 := 2
	if result := search(nums6, target6); result != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result)
	}

	// Test case 7: The array contains just one element and the target is present
	nums7 := []int{1}
	target7 := 1
	expected7 := 0
	if result := search(nums7, target7); result != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result)
	}

	// Test case 8: The array contains just one element and the target is not present
	nums8 := []int{1}
	target8 := 2
	expected8 := -1
	if result := search(nums8, target8); result != expected8 {
		t.Errorf("Test case 8 failed: expected %d, got %d", expected8, result)
	}

	// Test case 9: The array is rotated at the first element
	nums9 := []int{2, 3, 4, 5, 6, 7, 1}
	target9 := 1
	expected9 := 6
	if result := search(nums9, target9); result != expected9 {
		t.Errorf("Test case 9 failed: expected %d, got %d", expected9, result)
	}

	// Test case 10:
	nums10 := []int{3, 5, 1}
	target10 := 3
	expected10 := 0
	if result := search(nums10, target10); result != expected10 {
		t.Errorf("Test case 10 failed: expected %d, got %d", expected10, result)
	}

	// Test case 11:
	nums11 := []int{3, 1}
	target11 := 3
	expected11 := 0
	if result := search(nums11, target11); result != expected11 {
		t.Errorf("Test case 11 failed: expected %d, got %d", expected11, result)
	}

	// Test case 12:
	nums12 := []int{5, 1, 3}
	target12 := 5
	expected12 := 0
	if result := search(nums12, target12); result != expected12 {
		t.Errorf("Test case 12 failed: expected %d, got %d", expected12, result)
	}
}
