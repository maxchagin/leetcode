package finalarraystateafterkmultiplicationoperationsi

func getFinalState(nums []int, k int, multiplier int) []int {

	for c := 0; c < k; c++ {
		min := 0
		for i := range nums {
			if nums[i] < nums[min] {
				min = i
			}
		}
		nums[min] *= multiplier
	}

	return nums
}
