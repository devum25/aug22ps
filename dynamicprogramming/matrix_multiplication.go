package dynamicprogramming

import "math"

// Given an array of integers A representing chain of 2-D matices such that the dimensions of ith matrix is A[i-1] x A[i].
// Find the most efficient way to multiply these matrices together. The problem is not actually to perform the multiplications,
// but merely to decide in which order to perform the multiplications.
// Return the minimum number of multiplications needed to multiply the chain.

func MinMatrixMultiplication(A []int) int {
	row := make([]int, len(A)-1)
	col := make([]int, len(A)-1)

	for i := 0; i < len(A)-1; i++ {
		row[i] = A[i]
		col[i] = A[i+1]
	}

	col[len(col)-1] = A[len(A)-1]

	dp := make([][]int, len(row)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(row)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	x := multiply(row, col, 0, len(row)-1, dp)
	return x
}

func multiply(row []int, col []int, s, e int, dp [][]int) int {
	if s == e {
		// dp[s][e] = 0
		return 0
	}
	if dp[s][e] == -1 {
		ans := math.MaxInt
		for i := s; i < e; i++ {
			x := multiply(row, col, s, i, dp) + multiply(row, col, i+1, e, dp) + row[s]*col[e]*col[i]
			ans = min(ans, x)
		}
		dp[s][e] = ans
	}

	return dp[s][e]
}
