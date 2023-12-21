package dynamicprogramming

// 198. House Robber
// Companies
// You are a professional robber planning to rob houses along a street.
//  Each house has a certain amount of money stashed,
//  the only constraint stopping you from robbing each of them is that
//  adjacent houses have security systems connected and it will automatically
//  contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 2:

// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.

func Rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	x := money(nums, &dp, len(nums)-1)
	return x
}

func money(n []int, dp *[]int, i int) int {
	if i == 0 {
		return n[i]
	}
	if i == 1 {
		return max(n[i], n[0])
	}

	if (*dp)[i] != -1 {
		return (*dp)[i]
	}

	(*dp)[i] = max(n[i]+money(n, dp, i-2), money(n, dp, i-1))

	return (*dp)[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
