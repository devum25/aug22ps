package dynamicprogramming

// You are given a set of coins A, In how many ways you can make sum B.
// You can choose one coin exactly once.

func CoinChange_0_1(A []int, B int) int {
	dp := make([][]int, len(A)+1)
	n := B + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	r := 0
	for r < len(dp) {
		dp[r][0] = 1
		r++
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j]

			if j >= A[i-1] {
				dp[i][j] = dp[i-1][j-A[i-1]]
			}
		}
	}

	return dp[len(A)][B]
}
