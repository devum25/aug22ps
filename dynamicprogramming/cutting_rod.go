package dynamicprogramming

func CuttingRod(A []int) int {
	dp := make([][]int, len(A)+1)
	k := len(A)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return KnapSackRod(A, dp, k, k)
}

func KnapSackRod(A []int, dp [][]int, k int, n int) int {
	if dp[n][k] == -1 {
		if n == 0 || k == 0 {
			dp[n][k] = 0
		} else {
			ans := KnapSackRod(A, dp, k, n-1)
			if k >= n {
				dp[n][k] = max(ans, KnapSackRod(A, dp, k-n, n-1)+A[n-1])
			} else {
				dp[n][k] = ans
			}
		}
	}

	return dp[n][k]
}
