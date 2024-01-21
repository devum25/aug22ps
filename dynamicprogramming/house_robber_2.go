package dynamicprogramming

// 213. House Robber II
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
// That means the first house is the neighbor of the last one.
//  Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

func Rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 0 {
		return 0
	}
	dp1 := make([]int, len(nums)-1)
	dp2 := make([]int, len(nums)-1)

	for i := 0; i < len(dp1); i++ {
		dp1[i] = -1
		dp2[i] = -1
	}
	x := robfisthouse(nums[0:len(nums)-1], dp1, len(nums)-2)
	y := roblasthouse(nums[1:], dp2, len(nums)-2)

	return max(x, y)
}

func robfisthouse(nums, dp []int, i int) int {
	if i == 0 {
		return nums[i]
	}
	if i < 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max(nums[i]+robfisthouse(nums, dp, i-2), robfisthouse(nums, dp, i-1))
	return dp[i]
}

func roblasthouse(nums, dp []int, i int) int {
	if i == 0 {
		return nums[i]
	}
	if i < 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max(nums[i]+roblasthouse(nums, dp, i-2), roblasthouse(nums, dp, i-1))
	return dp[i]
}
