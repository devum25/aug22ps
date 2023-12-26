package dynamicprogramming

// You are given an m x n integer array grid. There is a robot initially located at the top-left corner (i.e., grid[0][0]).
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
// The robot can only move either down or right at any point in time.
// An obstacle and space are marked as 1 or 0 respectively in grid.
// A path that the robot takes cannot include any square that is an obstacle.
// Return the number of possible unique paths that the robot can take to reach the bottom-right corner.
// The testcases are generated so that the answer will be less than or equal to 2 * 109.

// Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// Output: 2
// Explanation: There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right

/////////////////////////////////////////////////////////
func UniquePathsWithObstaclesIterative(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// Check if the starting or ending cell is an obstacle
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// Initialize the dp array
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: There is one way to reach the starting cell
	dp[0][0] = 1

	// Fill in the dp array
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Skip obstacles
			if obstacleGrid[i][j] == 1 {
				continue
			}

			// Update the number of paths from the left and above cells
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//////////////////////////////////////////////
// TOP DOWN APPROACH

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// Check if the starting or ending cell is an obstacle
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// Initialize the dp array with -1 to represent uncomputed values
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// Call the recursive function with memoization
	return overcomeObstacle(0, 0, m, n, obstacleGrid, dp)
}

func overcomeObstacle(i, j, m, n int, obstacle [][]int, dp [][]int) int {
	// Check if out of bounds or obstacle
	if i >= m || j >= n || obstacle[i][j] == 1 {
		return 0
	}

	// Check if reached the destination
	if i == m-1 && j == n-1 {
		return 1
	}

	// Check if already computed
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	// Recursive calls for moving right and down
	rightPaths := overcomeObstacle(i, j+1, m, n, obstacle, dp)
	downPaths := overcomeObstacle(i+1, j, m, n, obstacle, dp)

	// Update the dp array with the result for the current cell
	dp[i][j] = rightPaths + downPaths

	return dp[i][j]
}
