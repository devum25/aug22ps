package dynamicprogramming

import "math"

// 64. Minimum Path Sum
// Given a m x n grid filled with non-negative numbers
// find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.

// Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
// Output: 7
// Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

func MinPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	n := len(grid[0])

	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	x := minpath(0, 0, len(grid), n, grid, dp)

	return x
}

func minpath(i, j, m, n int, grid [][]int, dp [][]int) int {
	if i == m-1 && j == n-1 {
		return grid[i][j]
	}

	if (dp)[i][j] != -1 {
		return (dp)[i][j]
	}

	right := math.MaxInt32
	if j+1 < n {
		right = minpath(i, j+1, m, n, grid, dp)
	}

	down := math.MaxInt32
	if i+1 < m {
		down = minpath(i+1, j, m, n, grid, dp)
	}

	(dp)[i][j] = grid[i][j] + min(right, down)
	return (dp)[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/////////////////////////////////////////////////
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Create a memoization cache
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Define the recursive helper functio

	return dfs(0, 0, memo, grid, m, n)
}

func dfs(i, j int, memo [][]int, grid [][]int, m, n int) int {
	if i == m-1 && j == n-1 {
		return grid[i][j]
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	right := math.MaxInt32
	if j+1 < n {
		right = dfs(i, j+1, memo, grid, m, n)
	}

	down := math.MaxInt32
	if i+1 < m {
		down = dfs(i+1, j, memo, grid, m, n)
	}

	memo[i][j] = grid[i][j] + min(right, down)
	return memo[i][j]
}
