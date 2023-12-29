package dynamicprogramming

// 72. Edit Distance
// Medium
// 14.2K
// 189
// Companies
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
// You have the following three operations permitted on a word:
// Insert a character
// Delete a character
// Replace a character

// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')

func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	n := len(word2) + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}

	return minimum(len(word1), len(word2), word1, word2, dp)
}

func minimum(l1, l2 int, s1, s2 string, dp [][]int) int {
	if l2 == 0 {
		return l1
	}
	if l1 == 0 {
		return l2
	}
	if dp[l1][l2] != -1 {
		return dp[l1][l2]
	}
	if s1[l1-1] == s2[l2-1] {
		dp[l1][l2] = minimum(l1-1, l2-1, s1, s2, dp)
	} else {
		dp[l1][l2] = min1(1+minimum(l1-1, l2, s1, s2, dp), 1+minimum(l1, l2-1, s1, s2, dp), 1+minimum(l1-1, l2-1, s1, s2, dp))
	}

	return dp[l1][l2]
}
