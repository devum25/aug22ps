package arrayquestion

func majorityElement(nums []int) int {
	majority := nums[0]
	vote := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			vote++
		} else if nums[i] != majority {
			vote--
		}

		if vote == 0 {
			majority = nums[i]
			vote = 1
		}
	}

	return majority
}
