package dynamicprogramming

import "math"

// 279. Perfect Squares
// Companies
// Given an integer n, return the least number of perfect square numbers that sum to n.
// A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

// Example 1:

// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4.
// Example 2:

// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9.

func NumSquares(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	perfect(n, &dp)
	return dp[n]
}

func perfect(n int, dp *[]int) int {
	if n == 1 || n == 0 {
		(*dp)[n] = n
		return n
	}

	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	min := math.MaxInt

	for i := 1; i*i <= n; i++ {
		x := perfect(n-(i*i), dp) + 1
		min = int(math.Min(float64(min), float64(x)))
	}

	(*dp)[n] = min
	return (*dp)[n]
}
