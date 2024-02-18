package subarrays

import "math"

func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = 1 * (n + 1)
		}
	}

	for i := 0; i < len(nums); i++ {
		if int(math.Abs(float64(nums[i]))) >= 1 && int(math.Abs(float64(nums[i]))) <= n {
			idx := int(math.Abs(float64(nums[i]))) - 1
			nums[idx] = -1 * int(math.Abs(float64(nums[idx])))
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}
