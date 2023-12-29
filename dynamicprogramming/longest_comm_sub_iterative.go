package dynamicprogramming

func LongestCommonSubsequenceIterative(text1 string, text2 string) int {
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
