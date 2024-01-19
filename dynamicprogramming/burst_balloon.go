package dynamicprogramming

import "math"

// You are given n balloons, indexed from 0 to n - 1.
// Each balloon is painted with a number on it represented by an array nums. You are asked to burst all the balloons.
// If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins.
// If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.
// Return the maximum coins you can collect by bursting the balloons wisely.

// Code
// Testcase
// Test Result
// Test Result
// 312. Burst Balloons
// Hard
// Topics
// Companies
// You are given n balloons, indexed from 0 to n - 1. Each balloon is painted with a number on it represented by an array nums. You are asked to burst all the balloons.

// If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins. If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.

// Return the maximum coins you can collect by bursting the balloons wisely.

// Example 1:

// Input: nums = [3,1,5,8]
// Output: 167
// Explanation:
// nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
// Example 2:

// Input: nums = [1,5]
// Output: 10
func BurstBalloon(nums []int) int {
	n := len(nums)
	// Add boundary balloons with value 1 at the beginning and end
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	// Initialize a 2D DP array to store the maximum coins
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	return burst(nums, dp, 1, n)
}

func burst(nums []int, dp [][]int, left, right int) int {
	if left > right {
		return 0
	}
	if dp[left][right] > 0 {
		return dp[left][right]
	}

	maxCoins := 0
	for i := left; i <= right; i++ {
		coins := nums[left-1]*nums[i]*nums[right+1] +
			burst(nums, dp, left, i-1) +
			burst(nums, dp, i+1, right)

		maxCoins = int(math.Max(float64(maxCoins), float64(coins)))
	}

	dp[left][right] = maxCoins
	return maxCoins
}
