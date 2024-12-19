package linkedlistcycle

import "testing"

func createList(nodes []int, pos int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	head := &ListNode{Val: nodes[0]}
	current := head
	var cycleNode *ListNode
	for i := 1; i < len(nodes); i++ {
		current.Next = &ListNode{Val: nodes[i]}
		current = current.Next
		if i == pos {
			cycleNode = current
		}
	}
	if pos != -1 {
		current.Next = cycleNode
	}
	return head
}

func TestHasCycle(t *testing.T) {
	// Case 1: List with no elements
	head := createList([]int{}, -1)
	if hasCycle(head) != false {
		t.Errorf("Failed case 1: expected false, got true")
	}

	// Case 2: List with single element, no cycle
	head = createList([]int{1}, -1)
	if hasCycle(head) != false {
		t.Errorf("Failed case 2: expected false, got true")
	}

	// Case 3: List with two elements, no cycle
	head = createList([]int{1, 2}, -1)
	if hasCycle(head) != false {
		t.Errorf("Failed case 3: expected false, got true")
	}

	// Case 4: List with three elements, no cycle
	head = createList([]int{1, 2, 3}, -1)
	if hasCycle(head) != false {
		t.Errorf("Failed case 5: expected false, got true")
	}

	// Case 5: List with three elements, cycle at the second element
	head = createList([]int{1, 2, 3}, 1)
	if hasCycle(head) != true {
		t.Errorf("Failed case 6: expected true, got false")
	}

	// Case 6: Larger list with cycle in the middle
	head = createList([]int{1, 2, 3, 4, 5, 6}, 2)
	if hasCycle(head) != true {
		t.Errorf("Failed case 7: expected true, got false")
	}

	// Case 7: Larger list with cycle at the end
	head = createList([]int{1, 2, 3, 4, 5, 6}, 5)
	if hasCycle(head) != true {
		t.Errorf("Failed case 8: expected true, got false")
	}

	// Case 8: List with a cycle where pos is -1 (indicating no cycle)
	head = createList([]int{3, 2, 0, -4}, -1)
	if hasCycle(head) != false {
		t.Errorf("Failed case 9: expected false, got true")
	}

	// Case 9: List with five elements, cycle at the start
	head = createList([]int{1, 2, 3, 4, 5}, 0)
	if hasCycle(head) != false {
		t.Errorf("Failed case 9: expected true, got false")
	}

	// Case 10: List with five elements, cycle at the last element
	head = createList([]int{1, 2, 3, 4, 5}, 4)
	if hasCycle(head) != true {
		t.Errorf("Failed case 10: expected true, got false")
	}
}
