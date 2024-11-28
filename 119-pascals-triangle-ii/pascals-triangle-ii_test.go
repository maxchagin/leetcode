package pascalstriangleii

import (
	"fmt"
	"reflect"
	"testing"
)

// Test case structure for the getRow function
type testCase struct {
	rowIndex int
	expected []int
}

// TestGetRow tests the getRow function with various test cases.
func TestGetRow(t *testing.T) {
	tests := []testCase{
		// Test case 1: Base case with rowIndex 0
		{rowIndex: 0, expected: []int{1}},

		// Test case 2: First non-trivial row with rowIndex 1
		{rowIndex: 1, expected: []int{1, 1}},

		// Test case 3: rowIndex 2 forms an incremental pattern
		{rowIndex: 2, expected: []int{1, 2, 1}},

		// Test case 4: Known pattern with rowIndex 3
		{rowIndex: 3, expected: []int{1, 3, 3, 1}},

		// Test case 5: Check rowIndex 4
		{rowIndex: 4, expected: []int{1, 4, 6, 4, 1}},

		// Test case 6: Higher row where complexity starts to increase
		{rowIndex: 5, expected: []int{1, 5, 10, 10, 5, 1}},

		// Test case 7: Mid-size row index to test calculation correctness
		{rowIndex: 6, expected: []int{1, 6, 15, 20, 15, 6, 1}},

		// Test case 8: Boundary test for larger rowIndex 10
		{rowIndex: 10, expected: []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1}},

		// Test case 9: Further test to ensure handling of row computations
		{rowIndex: 7, expected: []int{1, 7, 21, 35, 35, 21, 7, 1}},

		// Test case 10: Another mid-size test case for accuracy
		{rowIndex: 8, expected: []int{1, 8, 28, 56, 70, 56, 28, 8, 1}},
	}

	for _, tt := range tests {
		// Use fmt.Sprintf to convert rowIndex to a string
		t.Run(fmt.Sprintf("rowIndex_%d", tt.rowIndex), func(t *testing.T) {
			result := getRow(tt.rowIndex)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("getRow(%d) = %v, expected %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}
