package dynamicprogramming

// 70. Climbing Stairs
// Easy
// 20.6K
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Example 1:

// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:

// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

// iterative way
func ClimbStairsIterative(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[n] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return climb(n, &dp)

}

func climb(n int, dp *[]int) int {
	if n <= 1 {
		return 1
	}
	if (*dp)[n] == -1 {
		(*dp)[n] = climb(n-1, dp) + climb(n-2, dp)
	}
	return (*dp)[n]
}
