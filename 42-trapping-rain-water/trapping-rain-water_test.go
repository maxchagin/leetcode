package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	var height []int
	var expected int
	var result int

	// // Test case 0: Easy example from the problem description.
	// height = []int{1, 0, 2, 1, 0, 1, 3}
	// expected = 5
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 0 failed. Expected %d, got %d", expected, result)
	// }
	// // Test case 1: Basic example from the problem description.
	// height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// expected = 6
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 1 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 2: Another example from the problem description.
	// height = []int{4, 2, 0, 3, 2, 5}
	// expected = 9
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 2 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 3: No bars.
	// // Should return zero, as there's no space to trap water.
	// height = []int{}
	// expected = 0
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 3 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 4: Single bar.
	// // Should return zero, as no water can be trapped.
	// height = []int{5}
	// expected = 0
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 4 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 5: Two bars.
	// // Should return zero, as water cannot be trapped between two bars.
	// height = []int{5, 5}
	// expected = 0
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 5 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 6: Uniform height.
	// // Should return zero, as there are no dips to trap water.
	// height = []int{3, 3, 3, 3, 3}
	// expected = 0
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 6 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 7: Peak and valleys.
	// height = []int{2, 0, 2}
	// expected = 2
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 7 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 8: Ascending and descending with dip.
	// height = []int{0, 3, 0, 2, 0, 4}
	// expected = 7
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 8 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 9: Large dip between two large bars.
	// height = []int{5, 1, 7}
	// expected = 4
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 9 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 10: Multiple small dips.
	// height = []int{5, 2, 1, 2, 1, 5}
	// expected = 14
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 10 failed. Expected %d, got %d", expected, result)
	// }

	// // Test case 11
	// height = []int{4, 2, 3}
	// expected = 1
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 11 failed. Expected %d, got %d", expected, result)
	// }
	// // Test case 12
	// height = []int{0, 7, 1, 4, 6}
	// expected = 7
	// result = trap(height)
	// if result != expected {
	// 	t.Errorf("Test case 11 failed. Expected %d, got %d", expected, result)
	// }
	// Test case 13
	height = []int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}
	expected = 26
	result = trap(height)
	if result != expected {
		t.Errorf("Test case 11 failed. Expected %d, got %d", expected, result)
	}
}
