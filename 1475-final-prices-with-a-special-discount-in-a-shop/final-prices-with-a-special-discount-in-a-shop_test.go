package finalpriceswithaspecialdiscountinashop

import (
	"reflect"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	// Test case 1: Basic scenario with significant discounts
	prices := []int{8, 4, 6, 2, 3}
	expected := []int{4, 2, 4, 2, 3}
	result := finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 2: No discounts for any items
	prices = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 3: Mixed discounts, including zero final prices
	prices = []int{10, 1, 1, 6}
	expected = []int{9, 0, 1, 6}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 3 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 4: Single item in the array
	prices = []int{7}
	expected = []int{7}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 4 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 5: All elements are the same
	prices = []int{5, 5, 5, 5, 5}
	expected = []int{0, 0, 0, 0, 5}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 5 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 6: Decreasing prices with no discounts
	prices = []int{10, 9, 8, 7, 6}
	expected = []int{1, 1, 1, 1, 6}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 6 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 7: Prices where each discount equals the next item's price
	prices = []int{9, 1, 8, 1, 7, 1}
	expected = []int{8, 0, 7, 0, 6, 1}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 7 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 8: Repeating numbers with no discounts
	prices = []int{4, 4, 4, 4, 4}
	expected = []int{0, 0, 0, 0, 4}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 8 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 9: All prices are the same
	prices = []int{2, 2, 2, 2, 2}
	expected = []int{0, 0, 0, 0, 2}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 9 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}

	// Test case 10: Prices with large variance
	prices = []int{100, 1, 99, 1, 98, 1}
	expected = []int{99, 0, 98, 0, 97, 1}
	result = finalPrices(prices)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 10 failed. Prices %v Expected %v, got %v", prices, expected, result)
	}
}
