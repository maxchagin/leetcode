package linkedlistcycleii

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// Helper function to create a linked list with a cycle
	createCycleList := func(values []int, pos int) *ListNode {
		if len(values) == 0 {
			return nil
		}

		head := &ListNode{Val: values[0]}
		current := head

		var cycleEntry *ListNode
		for i := 1; i < len(values); i++ {
			current.Next = &ListNode{Val: values[i]}
			current = current.Next
			if i == pos {
				cycleEntry = current
			}
		}

		if pos != -1 {
			current.Next = cycleEntry
		}

		return head
	}

	// Test cases
	tests := []struct {
		values []int
		pos    int
		name   string
	}{
		{[]int{3, 2, 0, -4}, 1, "Test with cycle (pos index 1)"},
		{[]int{1}, -1, "Test without cycle"},
		{[]int{1, 2, 3, 4, 5}, -1, "Test with longer list without cycle"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, "Test with longer list with cycle at pos 4"},
		{[]int{2, 2, 2, 2}, 2, "Test with identical values and cycle"},
		{[]int{}, -1, "Test with empty list"},
		{[]int{1, 1, 1, 1, 1}, 3, "Test with cycle in middle of identical values list"},
		{[]int{10, 20, 30, 40, 50}, 2, "Test with cycle in non-identical values list"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			head := createCycleList(test.values, test.pos)
			result := detectCycle(head)

			var expected *ListNode
			if test.pos != -1 {
				expected = head
				for i := 0; i < test.pos; i++ {
					expected = expected.Next
				}
			}

			if result != expected {
				t.Errorf("Expected node with value %v, but got %v", expected, result)
			}
		})
	}
}
