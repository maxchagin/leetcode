package swapnodesinpairs

import (
	"reflect"
	"testing"
)

// Helper function to create linked list from slice.
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice.
func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var nums []int
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		// Test case 1: Normal case with even number of nodes
		{input: []int{1, 2, 3, 4}, output: []int{2, 1, 4, 3}},
		// Test case 2: Empty list
		{input: []int{}, output: []int{}},
		// Test case 3: Single node
		{input: []int{1}, output: []int{1}},
		// Test case 4: Odd number of nodes
		{input: []int{1, 2, 3}, output: []int{2, 1, 3}},
		// Test case 5: Two nodes
		{input: []int{1, 2}, output: []int{2, 1}},
		// Test case 6: Three nodes
		{input: []int{4, 5, 6}, output: []int{5, 4, 6}},
		// Test case 7: Four nodes
		{input: []int{7, 8, 9, 10}, output: []int{8, 7, 10, 9}},
		// Test case 8: Repeated nodes
		{input: []int{1, 1, 2, 2}, output: []int{1, 1, 2, 2}},
		// Test case 9: Larger even list
		{input: []int{9, 8, 7, 6, 5, 4, 3, 2}, output: []int{8, 9, 6, 7, 4, 5, 2, 3}},
		// Test case 10: Larger odd list
		{input: []int{11, 22, 33, 44, 55, 66, 77}, output: []int{22, 11, 44, 33, 66, 55, 77}},
	}

	for _, test := range tests {
		// Create linked list from input slice.
		head := createList(test.input)
		// Execute swapPairs function.
		result := swapPairs(head)
		// Convert result linked list to slice.
		resultSlice := listToSlice(result)

		// Assert that the result matches the expected output.
		if !reflect.DeepEqual(resultSlice, test.output) {
			t.Errorf("For input %v, expected output %v but got %v", test.input, test.output, resultSlice)
		}
	}
}
