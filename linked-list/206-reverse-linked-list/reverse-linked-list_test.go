package reverselinkedlist

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

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Example 2",
			input:    []int{1, 2},
			expected: []int{2, 1},
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
			// A case with repeating elements to check if the list correctly reverses
			name:     "Repeating Elements",
			input:    []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
		{
			// A typical case with an odd number of elements
			name:     "Odd Number of Elements",
			input:    []int{10, 20, 30, 40, 50},
			expected: []int{50, 40, 30, 20, 10},
		},
		{
			// A case where all elements are the same
			name:     "All Same Elements",
			input:    []int{7, 7, 7},
			expected: []int{7, 7, 7},
		},
		{
			// A list with alternating repeating elements
			name:     "Alternating Elements",
			input:    []int{1, 2, 1, 2, 1, 2},
			expected: []int{2, 1, 2, 1, 2, 1},
		},
		{
			// A list containing negative integers
			name:     "Negative Integers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-4, -3, -2, -1},
		},
		{
			// A case with a list that begins with zero
			name:     "List Starting with Zero",
			input:    []int{0, 1, 2, 3},
			expected: []int{3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.input)
			expected := createLinkedList(tt.expected)
			reversed := reverseList(head)
			// Convert the result and expected linked lists back to slices for comparison
			if !reflect.DeepEqual(linkedListToSlice(reversed), linkedListToSlice(expected)) {
				t.Errorf("Expected %v, but got %v", tt.expected, linkedListToSlice(reversed))
			}
		})
	}
}
