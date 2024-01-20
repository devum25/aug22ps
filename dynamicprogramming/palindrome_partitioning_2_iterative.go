package dynamicprogramming

// ababbbabbababa
func PalindromePartitioningIterative(s string) int {

	return minCut(s)
}

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)

	// Initialize dp array with the maximum possible cuts
	for i := range dp {
		dp[i] = i - 1
	}

	// Iterate through each center to find palindromes
	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2

		for left >= 0 && right < n && s[left] == s[right] {
			dp[right+1] = min(dp[right+1], dp[left]+1)
			left--
			right++
		}
	}

	return dp[n]
}
