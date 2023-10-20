package recursion

func DecimalToB(n int, b int) int {
	if n == 0 {
		return 0
	}

	return 10*DecimalToB(n/b, b) + n%5
}
