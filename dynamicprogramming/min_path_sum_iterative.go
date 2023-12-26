package dynamicprogramming

// 64. Minimum Path Sum

// Given a m x n grid filled with non-negative numbers,
// find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.
func MinPathSumIterative(grid [][]int) int {

	dp := make([][]int, len(grid))
	n := len(grid[0])
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, n)
	}

	i := len(grid) - 1
	dp[len(grid)-1][n-1] = grid[len(grid)-1][n-1]
	for i >= 1 {
		dp[i-1][n-1] = dp[i][n-1] + grid[i-1][n-1]
		i--
	}
	j := n - 1
	i = len(grid) - 1
	for j >= 1 {
		dp[i][j-1] = dp[i][j] + grid[i][j-1]
		j--
	}

	for r := i - 1; r >= 0; r-- {
		for c := n - 2; c >= 0; c-- {
			dp[r][c] = grid[r][c] + min(dp[r+1][c], dp[r][c+1])
		}
	}

	return dp[0][0]
}
