package dynamicprogramming

import "math"

// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
// A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
// Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

// Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
// Output: 13
// Explanation: There are two falling paths with a minimum sum as shown.

func MinFallingPathSum(matrix [][]int) int {
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
		for j := 0; j < len(matrix[0]); j++ {
			if j > 0 && j < len(matrix[0])-1 {
				dp[i][j] = matrix[i][j] + min1(dp[i+1][j], dp[i+1][j+1], dp[i+1][j-1])
			} else if j == 0 && len(matrix[0]) > 1 {
				dp[i][j] = matrix[i][j] + min(dp[i+1][j], dp[i+1][j+1])
			} else if j == len(matrix)-1 {
				dp[i][j] = matrix[i][j] + min(dp[i+1][j-1], dp[i+1][j])
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

func min1(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
