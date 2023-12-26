package dynamicprogramming

import "math"

// Given a triangle array, return the minimum path sum from top to bottom.
// For each step, you may move to an adjacent number of the row below.
// More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

// Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// Output: 11
// Explanation: The triangle looks like:
//    2
//   3 4
//  6 5 7
// 4 1 8 3
// The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

func MinimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	// n := len(triangle[0])

	dp[0][0] = triangle[0][0]
	dp[1][0] = dp[0][0] + triangle[1][0]
	dp[1][1] = dp[0][0] + triangle[1][1]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j != 0 && j < len(triangle[i-1]) {
				dp[i][j] = triangle[i][j] + min(dp[i-1][j], dp[i-1][j-1])
			} else if j == len(triangle[i-1]) {
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			}
		}
	}

	min := math.MaxInt
	row := len(dp) - 1
	for j := 0; j < len(dp[len(dp)-1]); j++ {
		if dp[row][j] < min {
			min = dp[row][j]
		}
	}

	return min
}
