package dynamicprogramming

// 416. Partition Equal Subset Sum

// Given an integer array nums,
//  return true if you can partition the array into two subsets such that the sum of the elements
//  in both subsets is equal or false otherwise.
func CanPartition(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	targetSum := sum / 2

	if sum%2 != 0 {
		return false
	}

	dp := make([][]int, len(nums)+1)
	n := targetSum + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	r := 0
	for r < len(dp) {
		dp[r][0] = 1
		r++
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j]

			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			}
		}
	}

	if dp[len(nums)][targetSum] != 0 && dp[len(nums)][targetSum]%2 == 0 {
		return true
	} else {
		return false
	}
}
