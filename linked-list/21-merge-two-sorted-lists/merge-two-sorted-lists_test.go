package mergetwosortedlists

import (
	"testing"
)

// Helper function to create list from slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice for easier comparison
func listToSlice(head *ListNode) []int {
	var result []int
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		// Test case 1: Both lists non-empty, with overlapping elements
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		// Test case 2: Both lists empty
		{list1: []int{}, list2: []int{}, expected: []int{}},
		// Test case 3: First list empty, second list non-empty
		{list1: []int{}, list2: []int{0}, expected: []int{0}},
		// Test case 4: Second list empty, first list non-empty
		{list1: []int{0}, list2: []int{}, expected: []int{0}},
		// Test case 5: Both lists have single element, different values
		{list1: []int{1}, list2: []int{2}, expected: []int{1, 2}},
		// Test case 6: Both lists have single element, same values
		{list1: []int{1}, list2: []int{1}, expected: []int{1, 1}},
		// Test case 7: Both lists have multiple elements, no overlap
		{list1: []int{1, 2, 3}, list2: []int{4, 5, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		// Test case 8: Lists with alternating merge pattern
		{list1: []int{1, 3, 5}, list2: []int{2, 4, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		// Test case 9: One list is a subset of the other
		{list1: []int{1, 2, 3}, list2: []int{2}, expected: []int{1, 2, 2, 3}},
		// Test case 10: All elements are equal in both lists
		{list1: []int{2, 2, 2}, list2: []int{2, 2, 2}, expected: []int{2, 2, 2, 2, 2, 2}},
		//Test case 11
		{list1: []int{2}, list2: []int{1}, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			list1 := createList(tt.list1)
			list2 := createList(tt.list2)

			// Call the mergeTwoLists function
			mergedList := mergeTwoLists(list1, list2)

			// Convert result to slice for comparison
			result := listToSlice(mergedList)

			// Compare the result with the expected output
			if !equal(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare two slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
