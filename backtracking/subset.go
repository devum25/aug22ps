package backtracking

func Subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	sub := make([]int, 0)
	solve1(nums, sub, &ans, 0)

	return ans
}

func solve1(nums, sub []int, ans *[][]int, idx int) {
	if idx == len(nums) {
		x := make([]int, len(sub))
		for i := 0; i < len(sub); i++ {
			x[i] = sub[i]
		}
		*ans = append(*ans, x)
		return
	}

	sub = append(sub, nums[idx])

	solve1(nums, sub, ans, idx+1)

	sub = sub[:len(sub)-1]
	solve1(nums, sub, ans, idx+1)
}
