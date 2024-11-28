package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	// Test case 1: Regular case with varying heights.
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	result := maxArea(height)
	if result != expected {
		t.Errorf("Test case 1 failed. Expected %d, got %d", expected, result)
	}

	// Test case 2: Two lines of equal height.
	height = []int{1, 1}
	expected = 1
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 2 failed. Expected %d, got %d", expected, result)
	}

	// Test case 3: Multiple identical heights.
	height = []int{2, 2, 2, 2, 2}
	expected = 8
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 3 failed. Expected %d, got %d", expected, result)
	}

	// Test case 4: Increasing heights from left to right.
	height = []int{1, 2, 3, 4, 5}
	expected = 6
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 4 failed. Expected %d, got %d", expected, result)
	}

	// Test case 5: Decreasing heights from left to right.
	height = []int{5, 4, 3, 2, 1}
	expected = 6
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 5 failed. Expected %d, got %d", expected, result)
	}

	// Test case 6: Only two heights.
	height = []int{3, 7}
	expected = 3
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 6 failed. Expected %d, got %d", expected, result)
	}

	// Test case 7: Heights with peak in the middle.
	height = []int{1, 3, 2, 5, 4, 3, 2, 1}
	expected = 12
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 7 failed. Expected %d, got %d", expected, result)
	}

	// Test case 8: All zero heights.
	height = []int{0, 0, 0, 0, 0}
	expected = 0
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 8 failed. Expected %d, got %d", expected, result)
	}

	// Test case 9: Some zero heights.
	height = []int{0, 6, 2, 0, 5}
	expected = 15
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 9 failed. Expected %d, got %d", expected, result)
	}

	// Test case 10: Heights form a symmetrical pattern.
	height = []int{4, 3, 2, 1, 2, 3, 4}
	expected = 24
	result = maxArea(height)
	if result != expected {
		t.Errorf("Test case 10 failed. Expected %d, got %d", expected, result)
	}
}
