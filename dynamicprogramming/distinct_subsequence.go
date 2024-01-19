package dynamicprogramming

// 115. Distinct Subsequences
// Given two strings s and t, return the number of distinct subsequences of s which equals t.
// The test cases are generated so that the answer fits on a 32-bit signed integer.

// Example 1:
// Input: s = "rabbbit", t = "rabbit"
// Output: 3
// Explanation:
// As shown below, there are 3 ways you can generate "rabbit" from s.
// rabb_it
// rab_bit
// rab_bbit

func DistinctSubsequence(s string, t string) int {
	dp := make([][]int, len(s)+1)
	n := len(t) + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	row := len(dp) - 1
	col := 0

	for row >= 0 {
		dp[row][col] = 1
		row--
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func NumDistinct1(s string, t string) int {
	dp := make([][]int, len(s)+1)
	n := len(t) + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return distinct1(s, t, len(s), len(t), dp)
}

func distinct1(s, t string, ls, lt int, dp [][]int) int {
	if dp[ls][lt] != -1 {
		return dp[ls][lt]
	}
	if lt == 0 {
		dp[ls][lt] = 1
	}
	if ls == 0 {
		dp[ls][lt] = 0
	}
	if s[ls-1] != t[lt-1] {
		dp[ls][lt] = distinct1(s, t, ls-1, lt, dp)
	} else {
		dp[ls][lt] = distinct1(s, t, ls-1, lt-1, dp) + distinct1(s, t, ls-1, lt, dp)
	}

	return dp[ls][lt]
}
