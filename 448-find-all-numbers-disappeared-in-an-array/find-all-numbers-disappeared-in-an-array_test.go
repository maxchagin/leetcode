package findallnumbersdisappearedinanarray

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	// Test case 1: Basic case with missing numbers
	input1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected1 := []int{5, 6}
	// Explanation: The numbers 5 and 6 are missing from 1 to 8
	assertFindDisappearedNumbers(t, input1, expected1)

	// Test case 2: All numbers are the same
	input2 := []int{1, 1}
	expected2 := []int{2}
	// Explanation: The number 2 is missing
	assertFindDisappearedNumbers(t, input2, expected2)

	// Test case 3: No number is missing
	input3 := []int{1, 2, 3, 4, 5}
	expected3 := []int{}
	// Explanation: All numbers from 1 to 5 are present
	assertFindDisappearedNumbers(t, input3, expected3)

	// Test case 4: Single element, missing element is 1
	input4 := []int{2}
	expected4 := []int{1}
	// Explanation: Number 1 is missing as range is [1,2]
	assertFindDisappearedNumbers(t, input4, expected4)

	// Test case 5: Range of size 1
	input5 := []int{1}
	expected5 := []int{}
	// Explanation: No number is missing, range is [1]
	assertFindDisappearedNumbers(t, input5, expected5)

	// Test case 6: Large input where a couple of numbers are missing
	input6 := []int{10, 9, 8, 7, 6, 5, 3, 3, 2, 1}
	expected6 := []int{4}
	// Explanation: Number 4 is missing in the range 1 to 10
	assertFindDisappearedNumbers(t, input6, expected6)

	// Test case 7: All elements are the same and missing a large range
	input7 := []int{2, 2, 2, 2}
	expected7 := []int{1, 3, 4}
	// Explanation: Numbers 1, 3, and 4 are missing
	assertFindDisappearedNumbers(t, input7, expected7)

	// Test case 8: All numbers are present in the sequence
	input8 := []int{5, 4, 3, 2, 1}
	expected8 := []int{}
	// Explanation: All numbers from 1 to 5 are present
	assertFindDisappearedNumbers(t, input8, expected8)

	// Test case 9: Multiple missing numbers
	input9 := []int{3, 3, 3, 3, 5, 5, 5, 6}
	expected9 := []int{1, 2, 4, 7, 8}
	// Explanation: Numbers 1, 2, 4, 7, 8 are missing in range 1 to 8
	assertFindDisappearedNumbers(t, input9, expected9)

	// Test case 10: Large numbers at the start
	input10 := []int{7, 8, 9, 10, 1, 2, 2, 1}
	expected10 := []int{3, 4, 5, 6}
	// Explanation: Numbers 3, 4, 5, 6 are missing in range 1 to 10
	assertFindDisappearedNumbers(t, input10, expected10)
}

func assertFindDisappearedNumbers(t *testing.T, input []int, expected []int) {
	result := findDisappearedNumbers(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For input %v, expected %v but got %v", input, expected, result)
	}
}
