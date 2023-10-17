package recursion

func IsPowerOfFour(n int) bool {
	var x float64
	x = float64(n / 4)
	if x == 1 {
		return true
	} else if x > 1 && x < 2 {
		return false
	}

	return IsPowerOfFour(n / 4)
}
