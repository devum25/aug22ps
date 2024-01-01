package dynamicprogramming

// 583. Delete Operation for Two Strings
// Given two strings word1 and word2,
// return the minimum number of steps required to make word1 and word2 the same.
// In one step, you can delete exactly one character in either string.

// Example 1:

// Input: word1 = "sea", word2 = "eat"
// Output: 2
// Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".

func MinDistanceDelete(word1 string, word2 string) int {
	lcs := longestCommonSubsequenceIterative(word1, word2)
	d1 := len(word1) - lcs
	d2 := len(word2) - lcs
	return d1 + d2
}

func longestCommonSubsequenceIterative(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	n := len(text2) + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
