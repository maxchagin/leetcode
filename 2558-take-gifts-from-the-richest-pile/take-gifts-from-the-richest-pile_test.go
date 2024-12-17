package takegiftsfromtherichestpile

import "testing"

func TestPickGifts(t *testing.T) {

	tests := []struct {
		gifts  []int // Input array of gifts
		k      int   // Number of seconds
		output int64 // Expected output
	}{
		// Test case 1: Basic scenario with mixed numbers
		{[]int{25, 64, 9, 4, 100}, 4, 29},

		// Test case 2: All piles are equal
		{[]int{1, 1, 1, 1}, 4, 4},

		// Test case 3: Single pile of gifts
		{[]int{100}, 1, 10},

		// Test case 4: Single pile of gifts, multiple seconds
		{[]int{100}, 3, 1},

		// Test case 5: Large number of gifts and pile
		{[]int{1000, 1000, 1000}, 5, 41},

		// Test case 6: Minimum size and values for gifts array
		{[]int{1}, 1, 1},

		// Test case 7: Maximum possible values in gifts array (but with reasonable k)
		{[]int{1000000000, 999999999, 999999998}, 3, 94866},

		// Test case 8: Large k value but not enough gifts to frequently pick from
		{[]int{2, 3, 1}, 10, 3},

		// Test case 9: Maximum possible k with small gifts
		{[]int{1, 2, 1, 1}, 1000, 4},

		// Test case 10: Random large values and large k
		{[]int{27, 81, 64, 16}, 6, 14},
	}

	for _, test := range tests {
		gifts := make([]int, len(test.gifts))
		copy(gifts, test.gifts)
		result := pickGifts(test.gifts, test.k)
		// Assert the correctness of the output
		if result != test.output {
			t.Errorf("For input gifts = %v and k = %d; expected %d but got %d", gifts, test.k, test.output, result)
		}
	}
}
