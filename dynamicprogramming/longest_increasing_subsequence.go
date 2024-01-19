package dynamicprogramming

// Given an integer array nums, return the length of the longest strictly increasing
// subsequence
// Example 1:

// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = 1
	res := 0
	for i := 1; i < len(nums); i++ {
		ans := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				ans = max(dp[j], ans)
			}
		}
		dp[i] = ans + 1
		res = max(res, dp[i])
	}

	return max(res, dp[0])
}
