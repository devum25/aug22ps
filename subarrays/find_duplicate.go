package subarrays

import "math"

func FindDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if int(math.Abs(float64(nums[i]))) >= 1 && int(math.Abs(float64(nums[i]))) <= n {
			idx := int(math.Abs(float64(nums[i]))) - 1
			if nums[idx] < 0 {
				return int(math.Abs(float64(nums[i])))
			}
			nums[idx] = -1 * nums[idx]
		}
	}

	return 0
}
