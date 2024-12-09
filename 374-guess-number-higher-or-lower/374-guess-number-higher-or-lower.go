package guessnumberhigherorlower

var pick int

func guess(num int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	} else {
		return 0
	}
}

// https://leetcode.com/problems/guess-number-higher-or-lower/
func guessNumber(n int) int {

	return -1
}
