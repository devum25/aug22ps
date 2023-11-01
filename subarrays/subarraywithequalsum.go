package subarrays

func FindSubarraysOfLengthWithEqualSum(nums []int) bool {

	hash := make(map[int]bool)

	for i := 0; i < len(nums)-1; i++ {
		x := nums[i] + nums[i+1]
		if _, ok := hash[x]; !ok {
			hash[x] = true
		} else if hash[x] {
			return true
		}
	}

	return false
}

// 1, 2, 3, 4, 5  --> 1,
