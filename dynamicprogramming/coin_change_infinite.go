package dynamicprogramming

// You are given a set of coins A. In how many ways can you make sum B assuming you have infinite amount of each coin in the set.
// NOTE:
// Coins in set A will be unique. Expected space complexity of this problem is O(B).
// The answer can overflow. So, return the answer % (106 + 7).
// https://www.youtube.com/watch?v=DJ4a7cmjZY0
func KnapSackCoinChange(A []int, B int) int {
	dp := make([][]int, len(A)+1)
	n := B + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	i := 0
	for i < len(dp) {
		dp[i][0] = 1
		i++
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if j >= A[i-1] {
				dp[i][j] = dp[i-1][j-A[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(A)][B]
}

// func coin(A []int, dp [][]int, k int, n int) int {
// 	if dp[n][k] == -1 {
// 		if n == 0 || k == 0 {
// 			dp[n][k] = 1
// 		}

// 		for i := 0; i < n; i++ {
// 			ans := coin(A, dp, k, n-1)
// 			if k >= A[i] {
// 				ans += coin(A, dp, k-A[i], n-1) + 1
// 			}
// 			dp[i][k] = ans
// 		}
// 	}

// 	return dp[n][k]
// }
