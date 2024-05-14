package dynamicprogramming

import "math"

func MaximumEnergy(nums []int, k int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MinInt
	}
	x := money1(nums, &dp, len(nums)-1, k)
	return x
}

func money1(n []int, dp *[]int, i int, k int) int {
	if i < 0 {
		return 0
	}
	if i == 0 {
		return n[i]
	}
	if (i - k) < 0 {
		return n[i]
	}

	if (*dp)[i] != math.MinInt {
		return (*dp)[i]
	}

	(*dp)[i] = max2(n[i], n[i]+n[i-k], money1(n, dp, i-1, k))

	return (*dp)[i]
}

func max2(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
