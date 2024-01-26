package dynamicprogramming

func TargetSum(n int, k int, target int) int {
	dp := make([][]int, target+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	x := roll1(target, k, n, dp)
	return x
}

func roll1(target int, k int, n int, dp [][]int) int {
	MOD := 1000000007
	if target == 0 && n != 0 {
		return 0
	}
	if target == 0 && n == 0 {
		return 1
	}
	if n <= 0 {
		return 0
	}

	if dp[target][n] != -1 {
		return dp[target][n]
	}

	ways := 0
	for i := 1; i <= k; i++ {
		if target-i >= 0 {
			ways = (ways + roll1(target-i, k, n-1, dp)) % MOD
		}
	}

	dp[target][n] = ways % MOD
	return dp[target][n]
}
