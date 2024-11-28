package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	// Test case 1: Basic example
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m1 := 3
	nums2 := []int{2, 5, 6}
	n1 := 3
	expected1 := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m1, nums2, n1)
	if !reflect.DeepEqual(nums1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, nums1)
	}

	// Test case 2: nums2 is empty
	nums2a := []int{1}
	m2 := 1
	nums2b := []int{}
	n2 := 0
	expected2 := []int{1}
	merge(nums2a, m2, nums2b, n2)
	if !reflect.DeepEqual(nums2a, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, nums2a)
	}

	// Test case 3: nums1 is empty
	nums3a := []int{0}
	m3 := 0
	nums3b := []int{1}
	n3 := 1
	expected3 := []int{1}
	merge(nums3a, m3, nums3b, n3)
	if !reflect.DeepEqual(nums3a, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, nums3a)
	}

	// Test case 4: Both arrays are empty
	nums4a := []int{0}
	m4 := 0
	nums4b := []int{}
	n4 := 0
	expected4 := []int{0}
	merge(nums4a, m4, nums4b, n4)
	if !reflect.DeepEqual(nums4a, expected4) {
		t.Errorf("Test case 4 failed. Expected %v, got %v", expected4, nums4a)
	}

	// Test case 5: Arrays are the same length
	nums5a := []int{1, 3, 5, 0, 0, 0}
	m5 := 3
	nums5b := []int{2, 4, 6}
	n5 := 3
	expected5 := []int{1, 2, 3, 4, 5, 6}
	merge(nums5a, m5, nums5b, n5)
	if !reflect.DeepEqual(nums5a, expected5) {
		t.Errorf("Test case 5 failed. Expected %v, got %v", expected5, nums5a)
	}

	// Test case 6: nums1 elements are smaller than nums2
	nums6a := []int{1, 2, 3, 0, 0, 0}
	m6 := 3
	nums6b := []int{4, 5, 6}
	n6 := 3
	expected6 := []int{1, 2, 3, 4, 5, 6}
	merge(nums6a, m6, nums6b, n6)
	if !reflect.DeepEqual(nums6a, expected6) {
		t.Errorf("Test case 6 failed. Expected %v, got %v", expected6, nums6a)
	}

	// Test case 7: nums2 elements are smaller than nums1
	nums7a := []int{4, 5, 6, 0, 0, 0}
	m7 := 3
	nums7b := []int{1, 2, 3}
	n7 := 3
	expected7 := []int{1, 2, 3, 4, 5, 6}
	merge(nums7a, m7, nums7b, n7)
	if !reflect.DeepEqual(nums7a, expected7) {
		t.Errorf("Test case 7 failed. Expected %v, got %v", expected7, nums7a)
	}

	// Test case 8: nums1 has repeated elements
	nums8a := []int{1, 1, 1, 0, 0, 0}
	m8 := 3
	nums8b := []int{2, 2, 2}
	n8 := 3
	expected8 := []int{1, 1, 1, 2, 2, 2}
	merge(nums8a, m8, nums8b, n8)
	if !reflect.DeepEqual(nums8a, expected8) {
		t.Errorf("Test case 8 failed. Expected %v, got %v", expected8, nums8a)
	}

	// Test case 9: nums2 has repeated elements
	nums9a := []int{2, 2, 2, 0, 0, 0}
	m9 := 3
	nums9b := []int{1, 1, 1}
	n9 := 3
	expected9 := []int{1, 1, 1, 2, 2, 2}
	merge(nums9a, m9, nums9b, n9)
	if !reflect.DeepEqual(nums9a, expected9) {
		t.Errorf("Test case 9 failed. Expected %v, got %v", expected9, nums9a)
	}

	// Test case 10: Large numbers
	nums10a := []int{100000, 100001, 100002, 0, 0, 0}
	m10 := 3
	nums10b := []int{99997, 99998, 99999}
	n10 := 3
	expected10 := []int{99997, 99998, 99999, 100000, 100001, 100002}
	merge(nums10a, m10, nums10b, n10)
	if !reflect.DeepEqual(nums10a, expected10) {
		t.Errorf("Test case 10 failed. Expected %v, got %v", expected10, nums10a)
	}
}
