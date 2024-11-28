package summaryranges

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	// Test case 1: Simple range with consecutive numbers.
	nums := []int{0, 1, 2, 4, 5, 7}
	expected := []string{"0->2", "4->5", "7"}
	result := summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Multiple individual numbers and ranges.
	nums = []int{0, 2, 3, 4, 6, 8, 9}
	expected = []string{"0", "2->4", "6", "8->9"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}

	// Test case 3: Single element.
	// Should return a single-element range.
	nums = []int{5}
	expected = []string{"5"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected, result)
	}

	// Test case 4: Empty array.
	// Should return an empty result.
	nums = []int{}
	expected = []string{}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 4 failed. Expected %v, got %v", expected, result)
	}

	// Test case 5: All elements make a single range.
	nums = []int{1, 2, 3, 4, 5}
	expected = []string{"1->5"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 5 failed. Expected %v, got %v", expected, result)
	}

	// Test case 6: No consecutive numbers.
	nums = []int{10, 20, 30}
	expected = []string{"10", "20", "30"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 6 failed. Expected %v, got %v", expected, result)
	}

	// Test case 7: Large range of numbers.
	nums = []int{-3, -2, -1, 0, 1, 2, 3, 4, 5}
	expected = []string{"-3->5"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 7 failed. Expected %v, got %v", expected, result)
	}

	// Test case 8: Negative numbers and ranges.
	nums = []int{-10, -9, -8, 0, 1, 2, 50}
	expected = []string{"-10->-8", "0->2", "50"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 8 failed. Expected %v, got %v", expected, result)
	}

	// Test case 9: Mixed single numbers and small ranges.
	nums = []int{-1, 1, 3, 5, 6, 8, 10}
	expected = []string{"-1", "1", "3", "5->6", "8", "10"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 9 failed. Expected %v, got %v", expected, result)
	}

	// Test case 10: Larger numbers array.
	nums = []int{100, 101, 102, 200, 300, 301, 302, 303}
	expected = []string{"100->102", "200", "300->303"}
	result = summaryRanges(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 10 failed. Expected %v, got %v", expected, result)
	}
}
