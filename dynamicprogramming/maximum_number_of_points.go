package dynamicprogramming

import "math"

// You are given an m x n integer matrix points (0-indexed). Starting with 0 points,
// you want to maximize the number of points you can get from the matrix.
// To gain points, you must pick one cell in each row.
// Picking the cell at coordinates (r, c) will add points[r][c] to your score.
// However, you will lose points if you pick a cell too far from the cell that you picked in the previous row.
// For every two adjacent rows r and r + 1 (where 0 <= r < m - 1), picking cells at coordinates (r, c1) and (r + 1, c2) will subtract abs(c1 - c2) from your score.
// Return the maximum number of points you can achieve.

// Input: points = [[1,2,3],[1,5,1],[3,1,1]]
// Output: 9
// Explanation:
// The blue cells denote the optimal cells to pick, which have coordinates (0, 2), (1, 1), and (2, 0).
// You add 3 + 5 + 3 = 11 to your score.
// However, you must subtract abs(2 - 1) + abs(1 - 0) = 2 from your score.
// Your final score is 11 - 2 = 9.
func MaxPoints(grid [][]int) int64 {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int64, len(grid))
	for i := range dp {
		dp[i] = make([]int64, len(grid[0]))
	}

	for i := 0; i < len(grid[0]); i++ {
		dp[0][i] = int64(grid[0][i])
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp := int64(math.MinInt64)
			for k := 0; k < len(grid[0]); k++ {
				dif := math.Abs(float64((k - j)))
				temp = max1(temp, dp[i-1][k]+int64(grid[i][j]-int(dif)))
			}
			dp[i][j] = temp
		}
	}

	ans := int64(math.MinInt64)
	for i := 0; i < len(grid[0]); i++ {
		ans = max1(ans, dp[len(grid)-1][i])
	}

	return ans
}

func max1(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
