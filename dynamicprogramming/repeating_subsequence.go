package dynamicprogramming

func RepeatingSubsequence(A string) int {
	dp := make([][]int, len(A)+1)

	n := len(A) + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if A[i-1] == A[j-1] && i != j {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	if dp[len(dp)-1][len(dp[0])-1] > 1 {
		return 1
	}

	return 0
}
