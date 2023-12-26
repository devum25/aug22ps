package dynamicprogramming

import "math"

// 1289. Minimum Falling Path Sum II
// Hard
// Companies
// Given an n x n integer matrix grid, return the minimum sum of a falling path with non-zero shifts.
// A falling path with non-zero shifts is a choice of exactly one element from each row of grid such that no two elements chosen in adjacent rows are in the same column.

// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: 13
// Explanation:
// The possible falling paths are:
// [1,5,9], [1,5,7], [1,6,7], [1,6,8],
// [2,4,8], [2,4,9], [2,6,7], [2,6,8],
// [3,4,8], [3,4,9], [3,5,7], [3,5,9]
// The falling path with the smallest sum is [1,5,7], so the answer is 13.

func MinFallingPathSum2(matrix [][]int) int {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0]
	}
	dp := make([][]int, len(matrix))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	row := len(matrix) - 1
	for c := 0; c < len(matrix[0]); c++ {
		dp[row][c] = matrix[row][c]
	}

	for i := row - 1; i >= 0; i-- {
		for j := i; j < len(matrix[0]); j++ {
			if j > 0 && j < len(matrix[0])-1 {
				dp[i][j-1] = matrix[i][j-1] + min1(dp[i+1][j], dp[i+1][j+1], dp[i+1][j-1])
				dp[i][j] = matrix[i][j] + min(dp[i+1][j-1], dp[i+1][j+1])
			} else if j == 0 && len(matrix[0]) > 1 {
				matrix[i][j] = matrix[i][j] + min(dp[i+1][j], dp[i+1][j+1])
				dp[i][j] = matrix[i][j] + dp[i+1][j+1]
			} else if j == len(matrix)-1 {
				matrix[i][j] = matrix[i][j] + min(dp[i+1][j-1], dp[i+1][j])
				dp[i][j] = matrix[i][j] + dp[i+1][j-1]
			}
		}
	}

	ans := math.MaxInt
	col := len(matrix[0])
	row = 0

	for c := 0; c < col; c++ {
		if dp[row][c] < ans {
			ans = dp[row][c]
		}
	}

	return ans
}
