package dynamicprogramming

func SellWine(nums []int) int {
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums)+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return sell(nums, 0, len(nums)-1, 1, dp)
}

func sell(nums []int, i, j, n int, dp [][]int) int {

	if i > j {
		return 0
	}

	if dp[i][j] == -1 {
		dp[i][j] = max(n*nums[i]+sell(nums, i+1, j, n+1, dp), n*nums[j]+sell(nums, i, j-1, n+1, dp))
	}

	return dp[i][j]
}
