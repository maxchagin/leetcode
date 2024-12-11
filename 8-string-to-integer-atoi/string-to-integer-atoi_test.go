package stringtointegeratoi

import "testing"

func TestMyAtoi(t *testing.T) {
	// Test case 1: Simple positive number
	input0 := "1234567890"
	expected0 := 1234567890
	if result := myAtoi(input0); result != expected0 {
		t.Errorf("Test case 0 failed: expected %d, got %d", expected0, result)
	}

	// Test case 1: Simple positive number
	input1 := "42"
	expected1 := 42
	if result := myAtoi(input1); result != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result)
	}

	// Test case 2: Simple negative number with leading spaces
	input2 := "   -42"
	expected2 := -42
	if result := myAtoi(input2); result != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result)
	}

	// Test case 3: Negative zero with leading zeros
	input3 := "-0012"
	expected3 := -12
	if result := myAtoi(input3); result != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result)
	}

	// Test case 4: Positive number with trailing non-digit characters
	input4 := "4193 with words"
	expected4 := 4193
	if result := myAtoi(input4); result != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result)
	}

	// Test case 5: Input starting with non-digit character
	input5 := "words and 987"
	expected5 := 0
	if result := myAtoi(input5); result != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result)
	}

	// Test case 6: Number exceeding the 32-bit positive range
	input6 := "91283472332"
	expected6 := 2147483647 // INT_MAX
	if result := myAtoi(input6); result != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result)
	}

	// Test case 7: Number exceeding the 32-bit negative range
	input7 := "-91283472332"
	expected7 := -2147483648 // INT_MIN
	if result := myAtoi(input7); result != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result)
	}

	// Test case 8: Input that starts with zero and a negative sign
	input8 := "0-1"
	expected8 := 0
	if result := myAtoi(input8); result != expected8 {
		t.Errorf("Test case 8 failed: expected %d, got %d", expected8, result)
	}

	// Test case 9: Mixed signs
	input9 := "+-12"
	expected9 := 0
	if result := myAtoi(input9); result != expected9 {
		t.Errorf("Test case 9 failed: expected %d, got %d", expected9, result)
	}

	// Test case 10: Number immediately followed by a non-digit character
	input10 := "3.14159"
	expected10 := 3
	if result := myAtoi(input10); result != expected10 {
		t.Errorf("Test case 10 failed: expected %d, got %d", expected10, result)
	}

	// Test case 11: Number +1
	input11 := "+1"
	expected11 := 1
	if result := myAtoi(input11); result != expected11 {
		t.Errorf("Test case 11 failed: expected %d, got %d", expected11, result)
	}

	// Test case 12
	input12 := "-+12"
	expected12 := 0
	if result := myAtoi(input12); result != expected12 {
		t.Errorf("Test case 12 failed: expected %d, got %d", expected12, result)
	}

	// Test case 13
	input13 := "   +0 123"
	expected13 := 0
	if result := myAtoi(input13); result != expected13 {
		t.Errorf("Test case 13 failed: expected %d, got %d", expected13, result)
	}
}
