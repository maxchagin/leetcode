package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	// Test case 1: Both arrays are empty
	// Edge case where both input arrays are empty, should handle gracefully.
	// The behavior in this case may depend on specifications, but we assume it's invalid input.
	nums1 := []int{}
	nums2 := []int{}
	expected := 0.0 // or some other convention for invalid input
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 1 failed. Expected %f, got %f", expected, result)
	}

	// Test case 2: One array is empty, the other has one element
	// Simple case, the median is the single element in nums2.
	nums1 = []int{}
	nums2 = []int{6}
	expected = 6.0
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 2 failed. Expected %f, got %f", expected, result)
	}

	// Test case 3: One array is empty, the other has multiple elements
	// Median is the middle element of nums2.
	nums1 = []int{}
	nums2 = []int{1, 2, 3, 4, 5}
	expected = 3.0
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 3 failed. Expected %f, got %f", expected, result)
	}

	// Test case 4: Both arrays have one element each
	// Median is the average of the two elements.
	nums1 = []int{1}
	nums2 = []int{2}
	expected = 1.5
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 4 failed. Expected %f, got %f", expected, result)
	}

	// Test case 5: Both arrays have equal size and distinct elements
	// Median is the average of the two middle elements after merging.
	nums1 = []int{1, 3}
	nums2 = []int{2, 4}
	expected = 2.5
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 5 failed. Expected %f, got %f", expected, result)
	}

	// Test case 6: Both arrays have different sizes
	// Check for the correct median element in an uneven merge.
	nums1 = []int{1, 3}
	nums2 = []int{2, 4, 5}
	expected = 3.0
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 6 failed. Expected %f, got %f", expected, result)
	}

	// Test case 7: Arrays contain negative numbers
	// Ensure the function handles negative numbers correctly.
	nums1 = []int{-5, -3, -1}
	nums2 = []int{-2, 0, 2}
	expected = -1.5
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 7 failed. Expected %f, got %f", expected, result)
	}

	// Test case 8: Arrays contain both negative and positive numbers
	nums1 = []int{-3, -1, 4}
	nums2 = []int{-2, 1, 5}
	expected = 0
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 8 failed. Expected %f, got %f", expected, result)
	}

	// Test case 9: Arrays with multiple identical elements
	// Both arrays include identical elements, should merge appropriately.
	nums1 = []int{1, 1, 1}
	nums2 = []int{1, 1, 1}
	expected = 1.0
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 9 failed. Expected %f, got %f", expected, result)
	}

	// Test case 10: Large arrays
	// Test performance and correctness with large input sizes.
	nums1 = make([]int, 1000)
	nums2 = make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums1[i] = 2*i + 1 // Odd numbers
		nums2[i] = 2*i + 2 // Even numbers
	}
	expected = 1000.5
	if result := findMedianSortedArrays(nums1, nums2); result != expected {
		t.Errorf("Test case 10 failed. Expected %f, got %f", expected, result)
	}
}
