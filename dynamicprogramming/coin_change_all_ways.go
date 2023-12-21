package dynamicprogramming

func CoinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	ans := 0
	ways(coins, amount, &dp, &ans)
	return ans
}

func ways(coins []int, n int, dp *[]int, ans *int) int {
	if n == 0 {
		*ans++
		// (*dp)[n] = *ans
		return 1
	}

	// if (*dp)[n] != -1 {
	// 	return (*dp)[n]
	// }

	for i := 0; i < len(coins); i++ {
		if n-coins[i] >= 0 {
			ways(coins, n-coins[i], dp, ans)
		}
	}

	return 0
}
