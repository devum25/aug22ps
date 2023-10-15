package recursion

func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}
