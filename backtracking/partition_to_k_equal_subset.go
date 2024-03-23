package backtracking

// https://www.youtube.com/watch?v=mBk4I0X46oI
func CanPartitionKSubsets(nums []int, k int) bool {
	target := 0

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	target = sum / k

	used := make([]bool, len(nums))

	return solve5(nums, 0, 0, target, k, used)
}

func solve5(nums []int, idx, subsetSum, target, k int, used []bool) bool {
	if k == 0 {
		return true
	}
	if subsetSum == target {
		return solve5(nums, 0, 0, target, k-1, used)
	}

	for i := idx; i < len(nums); i++ {
		if used[i] || nums[i]+subsetSum > target {
			continue
		}
		used[i] = true
		if solve5(nums, i+1, subsetSum+nums[i], target, k, used) {
			return true
		}
		used[i] = false
	}

	return false
}
