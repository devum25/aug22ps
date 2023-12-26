package dynamicprogramming

// 62. Unique Paths
// Medium
// 16K
// 421
// Companies
// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
// The robot can only move either down or right at any point in time.
// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
// The test cases are generated so that the answer will be less than or equal to 2 * 109.

// Example 1:
// Input: m = 3, n = 7
// Output: 28
// Example 2:
// Input: m = 3, n = 2
// Output: 3
// Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	i := 0
	for i < n {
		dp[m-1][i] = 1
		i++
	}
	j := 0
	for j < m {
		dp[j][n-1] = 1
		j++
	}
	return findpath(0, 0, m, n, &dp)
}

func findpath(i, j int, m, n int, dp *[][]int) int {
	if i >= m || j >= n {
		return 0
	}
	if i == m-1 || j == n-1 {
		return 1
	}

	if (*dp)[i][j] != 0 {
		return (*dp)[i][j]
	}

	(*dp)[i][j] = findpath(i+1, j, m, n, dp) + findpath(i, j+1, m, n, dp)

	return (*dp)[i][j]
}

////////////////////////////////////////////////////////////////////////////

func UniquePathIterative(m, n int) int {
	dp := make([][]int, m)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	i := 0
	for i < n {
		dp[m-1][i] = 1
		i++
	}
	j := 0
	for j < m {
		dp[j][n-1] = 1
		j++
	}

	for i := m - 1; i >= 1; i-- {
		for j := n - 1; j >= 1; j-- {
			dp[i-1][j-1] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[0][0]
}
