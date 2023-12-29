package dynamicprogramming

func NumDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	n := len(t) + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}

	return distinct(0, 0, dp, s, t)
}

func distinct(i, j int, dp [][]int, s, t string) int {
	if j == len(t) {
		return 1
	}
	if i == len(s) {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if s[i] == t[j] {
		dp[i][j] = distinct(i+1, j+1, dp, s, t) + distinct(i+1, j, dp, s, t)
	} else {
		dp[i][j] = distinct(i+1, j, dp, s, t)
	}

	return dp[i][j]
}
