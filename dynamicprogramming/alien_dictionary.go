package dynamicprogramming

func WaysAlienSignal(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}

	return waystosendsignal(n, n, dp)
}

func waystosendsignal(n int, i int, dp []int) int {
	if i == 1 {
		return 2
	} else if i == 2 {
		return 3
	}
	if i <= 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = waystosendsignal(n, i-1, dp) + waystosendsignal(n, i-2, dp)

	return dp[i]
}
