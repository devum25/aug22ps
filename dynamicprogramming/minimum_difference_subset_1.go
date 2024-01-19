package dynamicprogramming

// Given an integer array A containing N integers.
// You need to divide the array A into two subsets S1 and S2 such that the absolute difference between their sums is minimum.
// Find and return this minimum possible absolute difference.
// NOTE:
// Subsets can contain elements from A in any order (not necessary to be contiguous).
// Each element of A should belong to any one subset S1 or S2, not both.
// It may be possible that one subset remains empty.

func MinimumDifferenceSubset1(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	targetSum := sum / 2

	dp := make([][]int, len(nums)+1)
	k := targetSum + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k)
	}

	r := 0
	for r < len(dp) {
		dp[r][0] = 1
		r++
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			}
		}
	}

	if dp[len(nums)][targetSum] != 0 && dp[len(nums)][targetSum]%2 == 0 {
		if sum%2 == 0 {
			return 0
		} else {
			return 1
		}
	} else if dp[len(nums)][targetSum] > 0 {
		return (sum - 2*targetSum)
	} else {
		row := len(nums)
		col := targetSum
		for col >= 0 {
			if dp[row][col] > 0 {
				return (sum - 2*col)
			}
			col--
		}
	}

	return -1
}
