package sortlist

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
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

// Helper function to convert a linked list back to a slice
func linkedListToSlice(head *ListNode) []int {
	var nums []int
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}

func TestSortList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			input:    []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "Empty List",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single Element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			// A list already sorted in ascending order
			name:     "Already Sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			// A list sorted in descending order
			name:     "Reverse Order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			// A list with repeating elements
			name:     "Repeating Elements",
			input:    []int{2, 3, 2, 1, 3},
			expected: []int{1, 2, 2, 3, 3},
		},
		{
			// A list with all elements being the same
			name:     "All Elements Same",
			input:    []int{7, 7, 7, 7},
			expected: []int{7, 7, 7, 7},
		},
		{
			// A list containing negative integers
			name:     "Negative Integers",
			input:    []int{-3, -1, -2, -5, -4},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			// A list with alternating high and low values
			name:     "Alternating Values",
			input:    []int{10, 1, 8, 2, 6, 3},
			expected: []int{1, 2, 3, 6, 8, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			expected := createLinkedList(tt.expected)
			sorted := sortList(head)
			// Convert the result and expected linked lists back to slices for comparison
			if !reflect.DeepEqual(linkedListToSlice(sorted), linkedListToSlice(expected)) {
				t.Errorf("Expected %v, but got %v", tt.expected, linkedListToSlice(sorted))
			}
		})
	}
}
