package dynamicprogramming

func EditDistanceIterative(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	j := 1
	for j < len(dp[0]) {
		dp[0][j] = j
		j++
	}

	i := 1
	for i < len(dp) {
		dp[i][0] = i
		i++
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min1(dp[i-1][j]+1, dp[i-1][j-1]+1, dp[i][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
