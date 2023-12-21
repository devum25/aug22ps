package dynamicprogramming

import "math"

// 322. Coin Change
// Medium
// 18.1K
// 417
// Companies
// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
// Example 2:

// Input: coins = [2], amount = 3
// Output: -1
// Example 3:

// Input: coins = [1], amount = 0
// Output: 0

//  static void Main(string[] args){
// 	DynamicProgramming dp = new DynamicProgramming();
// 	int N = 15;
// 	int[] coins = { 1, 7, 10 };
// 	int[] dp1 = new int[100];
// 	int t = coins.Count();
// 	int res = dp.CoinChangeTopDownApproach(15, coins, t, dp1);
// 	Console.WriteLine(res);
//  }

// TOP DOWN APPROACH
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	x := solve(coins, amount, &dp)
	return x
}

func solve(coins []int, n int, dp *[]int) int {
	if n == 0 {
		return 0
	}
	if (*dp)[n] != -1 {
		return (*dp)[n]
	}
	min := math.MaxInt - 1
	for i := 0; i < len(coins); i++ {
		if n-coins[i] >= 0 {
			x := solve(coins, n-coins[i], dp)
			if x != -1 {
				min = int(math.Min(float64(min), float64(x+1)))
			}
		}
	}
	if min == math.MaxInt-1 {
		min = -1
	}
	(*dp)[n] = min
	return (*dp)[n]
}

// BOTTOM UP APPROACH
func BottomUpApproach(coins []int, m int) int {
	return minimumNumber(m, coins)
}
func minimumNumber(m int, coins []int) int {
	dp := make([]int, m+1)

	for i := 1; i <= m; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if (i-coins[j]) >= 0 && dp[i-coins[j]] != math.MaxInt {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
			}
		}
	}

	if dp[m] == math.MaxInt {
		return -1
	}

	return dp[m]
}
