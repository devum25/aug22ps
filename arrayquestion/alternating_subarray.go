package arrayquestion

func CountAlternatingSubarrays(nums []int) int64 {
	var sum int64
	sum = 1

	for i := 1; i < len(nums); i++ {
		prev := nums[i]
		k := i
		sum++
		for k > 0 {
			if nums[k-1] != prev {
				sum++
				prev = nums[k-1]
				k--
			} else {
				break
			}

		}
	}

	return sum
}
