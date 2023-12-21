package dynamicprogramming

func Fibonnaci(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return fib(n, &dp)
}

func fib(n int, dp *[]int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return n
	}

	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	(*dp)[n] = fib(n-1, dp) + fib(n-2, dp)

	return (*dp)[n]
}
