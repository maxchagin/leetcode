package arrangingcoins

import "testing"

func TestArrangeCoins(t *testing.T) {
	// Test case 1: Exactly enough coins for full rows
	n1 := 3
	expected1 := 2 // 1 + 2 = 3 complete rows
	if result := arrangeCoins(n1); result != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result)
	}

	// Test case 2: Not enough coins to complete the next row
	n2 := 5
	expected2 := 2 // The 3rd row is incomplete
	if result := arrangeCoins(n2); result != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result)
	}

	// Test case 3: Enough coins to complete additional rows
	n3 := 8
	expected3 := 3 // The 4th row is incomplete
	if result := arrangeCoins(n3); result != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result)
	}

	// Test case 4: One coin - only first row can be completed
	n4 := 1
	expected4 := 1
	if result := arrangeCoins(n4); result != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result)
	}

	// Test case 5: No coins - no complete row
	n5 := 0
	expected5 := 0
	if result := arrangeCoins(n5); result != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result)
	}

	// Test case 6: Large n, check efficiency and accuracy
	n6 := 1000
	expected6 := 44 // for n=1000, k=44 rows are complete (1+2+...+44 = 990, 45th row starts but incomplete)
	if result := arrangeCoins(n6); result != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result)
	}

	// Test case 7: Just enough coins to add an additional complete row
	n7 := 6
	expected7 := 3 // 1 + 2 + 3 = 6 complete rows
	if result := arrangeCoins(n7); result != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result)
	}

	// Test case 8: Prime number of coins
	n8 := 11
	expected8 := 4 // Because row 5 is incomplete with 1+2+3+4=10 coins used
	if result := arrangeCoins(n8); result != expected8 {
		t.Errorf("Test case 8 failed: expected %d, got %d", expected8, result)
	}

	// Test case 9: Large number of coins, check if program can handle the magnitude
	n9 := 1804289383
	// Expected output is calculated separately to avoid computational complexity here
	// for example purposes, assume a theoretical calculation of expected9.
	expected9 := 60070
	if result := arrangeCoins(n9); result != expected9 {
		t.Errorf("Test case 9 failed: expected %d, got %d", expected9, result)
	}

	// Test case 10: Minimal operational boundary when starting to form next row
	n10 := 15
	expected10 := 5 // Because row 5 is completed exactly with all coins used
	if result := arrangeCoins(n10); result != expected10 {
		t.Errorf("Test case 10 failed: expected %d, got %d", expected10, result)
	}
}
