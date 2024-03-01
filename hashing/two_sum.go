package hashset

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		if _, ok := hash[x]; ok {
			if hash[x] != i {
				ans = append(ans, i)
				ans = append(ans, hash[x])
				return ans
			}
		}
	}

	return ans
}
