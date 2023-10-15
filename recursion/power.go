package recursion

func Pow(a, n int) int {
	if n == 0 || a == 1 {
		return 1
	}

	halfPower := Pow(a, n/2)

	if n%2 == 0 {
		return halfPower * halfPower
	} else {
		return halfPower * halfPower * a
	}
}
