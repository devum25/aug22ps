package backtracking

func Subsets(nums []int) [][]int {
	curr := make([]int, 0)
	ans := make([][]int, 0)

	solve1(nums, &ans, &curr, 0)
	return ans
}

func solve1(nums []int, ans *[][]int, curr *[]int, i int) {
	if i == len(nums) { // base case
		temp := make([]int, len(*curr))
		x := *curr
		for i := 0; i < len(temp); i++ {
			temp[i] = x[i]
		}
		*ans = append(*ans, temp)
		return
	}
	y := *curr
	y = append(y, nums[i])
	*curr = y
	solve1(nums, ans, curr, i+1)
	x := *curr
	x = x[:len(x)-1]
	*curr = x
	solve1(nums, ans, curr, i+1)
}
