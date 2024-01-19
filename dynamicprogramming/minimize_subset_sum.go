package dynamicprogramming

import "math"

// 2035. Partition Array Into Two Arrays to Minimize Sum Difference

// You are given an integer array nums of 2 * n integers.
// You need to partition nums into two arrays of length n to minimize the absolute difference of the sums of the arrays.
// To partition nums, put each element of nums into one of the two arrays.
// Return the minimum possible absolute difference

// Input: nums = [3,9,7,3]
// Output: 2
// Explanation: One optimal partition is: [3,9] and [7,3].
// The absolute difference between the sums of the arrays is abs((3 + 9) - (7 + 3)) = 2.

func MinimumDifference(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	targetSum := sum / 2

	dp := make([][]int, len(nums)+1)
	k := int(math.Abs(float64(targetSum))) + 1

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
				dp[i][j] = dp[i-1][int(math.Abs(float64(j-nums[i-1])))] + dp[i-1][j]
			}
		}
	}

	if dp[len(nums)][targetSum] > 0 {
		return 0
	} else {
		row := len(nums)
		col := targetSum
		for col >= 0 {
			if dp[row][col] > 0 {
				return sum/2 - col
			}
			col--
		}
	}

	return -1
}
