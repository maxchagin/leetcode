package removeduplicatesfromsortedlistii

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	// Helper function to convert slice to linked list
	toLinkedList := func(nums []int) *ListNode {
		dummy := &ListNode{}
		current := dummy
		for _, num := range nums {
			current.Next = &ListNode{Val: num}
			current = current.Next
		}
		return dummy.Next
	}

	// Helper function to convert linked list to slice
	toSlice := func(head *ListNode) []int {
		if head == nil {
			return []int{}
		}
		var result []int
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	// Test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 3, 4, 4, 5}, expected: []int{1, 2, 5}},
		{input: []int{1, 1, 1, 2, 3}, expected: []int{2, 3}},
		{input: []int{1, 1, 2, 3, 3, 3, 4}, expected: []int{2, 4}},
		{input: []int{1, 2, 2, 3, 3, 4, 4}, expected: []int{1}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}}, // No duplicates
		{input: []int{1, 1, 1, 1, 1}, expected: []int{}},              // All duplicates
		{input: []int{2, 2, 3, 3, 4, 4, 5, 5}, expected: []int{}},     // All are duplicate pairs
		{input: []int{}, expected: []int{}},                           // Empty list
		{input: []int{1}, expected: []int{1}},                         // Single element list
		{input: []int{1, 1}, expected: []int{}},                       // Single duplicate pair
	}

	for _, test := range tests {
		head := toLinkedList(test.input)
		result := deleteDuplicates(head)
		if !reflect.DeepEqual(toSlice(result), test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, toSlice(result))
		}
	}
}
