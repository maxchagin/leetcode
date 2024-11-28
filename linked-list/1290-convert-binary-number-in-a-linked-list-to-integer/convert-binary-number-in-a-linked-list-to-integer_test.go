package convertbinarynumberinalinkedlisttointeger

import "testing"

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		input    *ListNode
		expected int
	}{
		{input: createLinkedList([]int{1, 0, 1}), expected: 5},
		{input: createLinkedList([]int{0}), expected: 0},
		{input: createLinkedList([]int{1, 0, 0, 1}), expected: 9},
		{input: createLinkedList([]int{1, 1, 1, 1}), expected: 15},
		{input: createLinkedList([]int{1}), expected: 1},
		{input: createLinkedList([]int{0, 0, 0}), expected: 0},
	}

	for _, test := range tests {
		if output := getDecimalValue(test.input); output != test.expected {
			t.Errorf("For input list %v, expected %d but got %d", test.input, test.expected, output)
		}
	}
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}
