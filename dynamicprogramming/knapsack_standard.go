package dynamicprogramming

func KnapSack(w []int, v []int, k int) int {
	dp := make([][]int, len(w)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return solveKnapsack(w, v, dp, len(w), k)
}

func solveKnapsack(w []int, v []int, dp [][]int, n, k int) int {
	if dp[n][k] == -1 {
		if n == 0 || k == 0 {
			dp[n][k] = 0
		} else {
			ans := solveKnapsack(w, v, dp, n-1, k)
			if k >= w[n-1] {
				ans = max(ans, solveKnapsack(w, v, dp, n-1, k-w[n-1])+v[n-1])
			}
			dp[n][k] = ans
		}
	}

	return dp[n][k]
}

//////////////////////////////////////////////////////////////////////////////////////////////////

func KnapSackIterative(w []int, v []int, k int) int {
	dp := make([][]int, len(w)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j]

			if j >= w[i-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}

	return dp[len(v)][k]
}

////////////////////////////////////////////////////////////////////////////
func KnapSackIterative_1D(w []int, v []int, k int) int {
	dp := make([]int, k+1)

	for i := 1; i <= len(w); i++ {
		for j := k; j >= w[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-w[i-1]]+v[i-1])
		}
	}

	return dp[k]
}
