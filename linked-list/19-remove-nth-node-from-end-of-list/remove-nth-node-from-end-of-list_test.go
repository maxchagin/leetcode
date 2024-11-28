package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func sliceToLinkedList(nums []int) *ListNode {
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

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var nums []int
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}

func TestRemoveNthFromEnd(t *testing.T) {
	// Test case 1: Basic example
	head1 := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	n1 := 2
	expected1 := []int{1, 2, 3, 5}
	result1 := linkedListToSlice(removeNthFromEnd(head1, n1))
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// // Test case 2: Remove the only element
	// head2 := sliceToLinkedList([]int{1})
	// n2 := 1
	// expected2 := []int{}
	// result2 := linkedListToSlice(removeNthFromEnd(head2, n2))
	// if !reflect.DeepEqual(result2, expected2) {
	// 	t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	// }

	// // Test case 3: Remove the last element
	// head3 := sliceToLinkedList([]int{1, 2})
	// n3 := 1
	// expected3 := []int{1}
	// result3 := linkedListToSlice(removeNthFromEnd(head3, n3))
	// if !reflect.DeepEqual(result3, expected3) {
	// 	t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	// }

	// // Test case 4: Remove the first element
	// head4 := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	// n4 := 5
	// expected4 := []int{2, 3, 4, 5}
	// result4 := linkedListToSlice(removeNthFromEnd(head4, n4))
	// if !reflect.DeepEqual(result4, expected4) {
	// 	t.Errorf("Test case 4 failed. Expected %v, got %v", expected4, result4)
	// }

	// // Test case 5: Remove middle element
	// head5 := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	// n5 := 3
	// expected5 := []int{1, 2, 4, 5}
	// result5 := linkedListToSlice(removeNthFromEnd(head5, n5))
	// if !reflect.DeepEqual(result5, expected5) {
	// 	t.Errorf("Test case 5 failed. Expected %v, got %v", expected5, result5)
	// }

	// // Test case 6: Linked list with two elements, remove first
	// head6 := sliceToLinkedList([]int{1, 2})
	// n6 := 2
	// expected6 := []int{2}
	// result6 := linkedListToSlice(removeNthFromEnd(head6, n6))
	// if !reflect.DeepEqual(result6, expected6) {
	// 	t.Errorf("Test case 6 failed. Expected %v, got %v", expected6, result6)
	// }

	// // Test case 7: Remove from linked list with duplicate values
	// head7 := sliceToLinkedList([]int{1, 2, 2, 3, 4})
	// n7 := 3
	// expected7 := []int{1, 2, 3, 4}
	// result7 := linkedListToSlice(removeNthFromEnd(head7, n7))
	// if !reflect.DeepEqual(result7, expected7) {
	// 	t.Errorf("Test case 7 failed. Expected %v, got %v", expected7, result7)
	// }

	// // Test case 8: All elements the same, remove one
	// head8 := sliceToLinkedList([]int{5, 5, 5, 5})
	// n8 := 2
	// expected8 := []int{5, 5, 5}
	// result8 := linkedListToSlice(removeNthFromEnd(head8, n8))
	// if !reflect.DeepEqual(result8, expected8) {
	// 	t.Errorf("Test case 8 failed. Expected %v, got %v", expected8, result8)
	// }

	// // Test case 9: Remove from longer list
	// head9 := sliceToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// n9 := 7
	// expected9 := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	// result9 := linkedListToSlice(removeNthFromEnd(head9, n9))
	// if !reflect.DeepEqual(result9, expected9) {
	// 	t.Errorf("Test case 9 failed. Expected %v, got %v", expected9, result9)
	// }

	// // Test case 10: Remove from list with varying negative and positive integers
	// head10 := sliceToLinkedList([]int{-1, 2, 3, -4, 5})
	// n10 := 4
	// expected10 := []int{-1, 3, -4, 5}
	// result10 := linkedListToSlice(removeNthFromEnd(head10, n10))
	// if !reflect.DeepEqual(result10, expected10) {
	// 	t.Errorf("Test case 10 failed. Expected %v, got %v", expected10, result10)
	// }
}
