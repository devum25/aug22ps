package dynamicprogramming

// 1143. Longest Common Subsequence
// Medium
// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.

// Example 1:
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
// Example 2:
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
// Example 3:
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.

func LongestCommonSubsequence(text1 string, text2 string) int {

	dp := make([][]int, len(text1)+1)
	n := len(text2) + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return lca(text1, text2, len(text1), len(text2), dp)
}

func lca(s1, s2 string, l1, l2 int, dp [][]int) int {
	if dp[l1][l2] == -1 {
		if l1 == 0 || l2 == 0 {
			dp[l1][l2] = 0
		} else if s1[l1-1] == s2[l2-1] {
			dp[l1][l2] = 1 + lca(s1, s2, l1-1, l2-1, dp)
		} else {

			dp[l1][l2] = max(lca(s1, s2, l1-1, l2, dp), lca(s1, s2, l1, l2-1, dp))
		}
	}

	return dp[l1][l2]
}
