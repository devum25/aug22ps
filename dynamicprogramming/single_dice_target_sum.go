package dynamicprogramming

import "math"

// Give a dice (6-faced) and a number N. Count the number of ways to get a sum N, if you can roll the dice
// as many times as required.
// N = 4
// [[4],[2,2],[2,1,1],[3,1],[1,1,1,1],[1,1,2],[1,2,1],[1,3]]
// ans = 8

// arr := [1,2,3,4,5,6]
func solveRollDice(A int) int {
	dp := make([]int, A+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	x := int(math.Pow(10, 9) + 7)
	y := roll(A, &dp)
	return int(y % x)

}

func roll(n int, dp *[]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}

	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	(*dp)[n] = roll(n-1, dp) + roll(n-2, dp) + roll(n-3, dp) + roll(n-4, dp) + roll(n-5, dp) + roll(n-6, dp)

	return (*dp)[n]
}

////////////////////////////////////////////////////////////////////////

// solve optimal
func solveOptimal(A int) int {
	const MOD = 1000000007
	dp := make([]int, A+1)
	dp[0] = 1

	for i := 1; i <= A; i++ {
		for j := 1; j <= 6; j++ {
			if i-j >= 0 {
				dp[i] = (dp[i] + dp[i-j]) % MOD
			}
		}
	}

	return dp[A]

}
