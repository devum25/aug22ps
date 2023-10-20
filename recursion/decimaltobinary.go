package recursion

func Decimaltobinary(n int) int {
	if n < 2 {
		return n
	}

	return 10*Decimaltobinary(n/2) + n%2
}
