package addtwonumbers

import (
	"reflect"
	"testing"
)

// Helper function to convert a slice to a linked list
func sliceToList(nums []int) *ListNode {
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

// Helper function to convert a linked list to a slice
func listToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

// Test cases for the addTwoNumbers function
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		// Test case 1: Both lists have multiple digits, no carry operation results further
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},

		// Test case 2: Both numbers are zeros
		{[]int{0}, []int{0}, []int{0}},

		// Test case 3: Different list sizes with carries leading to increased size
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},

		// Test case 4: One list is zero, the other is a single nonzero digit
		{[]int{0}, []int{1}, []int{1}},

		// Test case 5: Carry occurs at the end, creating an additional node
		{[]int{9}, []int{9}, []int{8, 1}},

		// Test case 6: Large number due to identical list sizes with carry
		{[]int{1, 8}, []int{0}, []int{1, 8}},

		// Test case 7: Simple case where carries are needed across nodes
		{[]int{2, 4}, []int{5, 6, 4}, []int{7, 0, 5}},

		// Test case 8: Complex carry situation across entire list
		{[]int{1}, []int{9, 9, 9, 9}, []int{0, 0, 0, 0, 1}},

		// Test case 9: Two numbers consist of only one digit with simple addition
		{[]int{5}, []int{7}, []int{2, 1}},

		// Test case 10: Test where the second list is single digit leading to carry
		{[]int{9, 9, 1, 9}, []int{1}, []int{0, 0, 2, 9}},
		// Test case 11
		{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, []int{5, 6, 4}, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		l1 := sliceToList(test.l1)
		l2 := sliceToList(test.l2)
		expected := sliceToList(test.expected)

		result := addTwoNumbers(l1, l2)
		resultSlice := listToSlice(result)
		expectedSlice := listToSlice(expected)

		if !reflect.DeepEqual(resultSlice, expectedSlice) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", test.l1, test.l2, resultSlice, expectedSlice)
		}
	}
}
