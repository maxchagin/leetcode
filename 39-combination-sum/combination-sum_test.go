package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// Helper function to check if two slices of slices are equivalent.
	// This does not consider the order of combinations.
	isEqual := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}
		for _, v := range a {
			found := false
			for _, w := range b {
				if reflect.DeepEqual(v, w) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	// Test case 1: Simple case with a single solution.
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{{7}, {2, 2, 3}}
	result := combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Multiple solutions with different combinations.
	candidates = []int{2, 3, 5}
	target = 8
	expected = [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}

	// Test case 3: Target cannot be formed.
	candidates = []int{2}
	target = 1
	expected = [][]int{}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected, result)
	}

	// Test case 4: Single candidate repeated.
	candidates = []int{1}
	target = 4
	expected = [][]int{{1, 1, 1, 1}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 4 failed. Expected %v, got %v", expected, result)
	}

	// Test case 5: Large target with multiple combinations.
	candidates = []int{3, 4, 5}
	target = 11
	expected = [][]int{{3, 3, 5}, {4, 4, 3}, {5, 3, 3}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 5 failed. Expected %v, got %v", expected, result)
	}

	// Test case 6: Larger numbers with limited combinations.
	candidates = []int{7, 11, 13}
	target = 18
	expected = [][]int{{7, 11}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 6 failed. Expected %v, got %v", expected, result)
	}

	// Test case 7: No combination of the numbers can form the target.
	candidates = []int{5, 10}
	target = 7
	expected = [][]int{}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 7 failed. Expected %v, got %v", expected, result)
	}

	// Test case 8: Target is zero, expect empty combination.
	candidates = []int{1, 2, 3}
	target = 0
	expected = [][]int{{}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 8 failed. Expected %v, got %v", expected, result)
	}

	// Test case 9: Candidates contain only one element equal to the target.
	candidates = []int{11}
	target = 11
	expected = [][]int{{11}}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 9 failed. Expected %v, got %v", expected, result)
	}

	// Test case 10: All candidates are larger than the target.
	candidates = []int{9, 10, 11}
	target = 8
	expected = [][]int{}
	result = combinationSum(candidates, target)
	if !isEqual(result, expected) {
		t.Errorf("Test case 10 failed. Expected %v, got %v", expected, result)
	}
}
