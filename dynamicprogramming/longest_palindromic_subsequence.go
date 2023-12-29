package dynamicprogramming

// 516. Longest Palindromic Subsequence
// Given a string s, find the longest palindromic subsequences length in s.
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: s = "bbbab"
// Output: 4
// Explanation: One possible longest palindromic subsequence is "bbbb".

func LongestPalindromeSubseq(A string) int {
	dp := make([][]int, len(A)+1)

	n := len(A) + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	s1 := A
	s2 := reverse(A)

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
