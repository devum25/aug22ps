package arrayquestion

import "math"

func MinimumDeletions(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	minIdx := 0
	maxIdx := 0
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}

		if nums[i] < min {
			min = nums[i]
			minIdx = i
		}
	}

	mid := (len(nums) - 1) / 2
	count := 0
	if minIdx < mid {
		count = minIdx
	} else if minIdx > mid {
		count = (len(nums) - minIdx)
	}

	if maxIdx < mid {
		count += maxIdx
	} else if maxIdx > mid {
		count += (len(nums) - maxIdx)
	}

	return count
}
