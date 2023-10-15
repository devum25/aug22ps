package recursion

import "math"

func MyPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}
	var halfPower float64

	halfPower = MyPow(x, n/2)
	// if n < 0 {
	// 	halfPower = 1 / halfPower
	// }
	y := int(math.Abs(float64(n)))
	if y%2 == 0 {
		return halfPower * halfPower
	} else {
		if n < 0 {
			return halfPower * halfPower * 1 / x
		}

		return halfPower * halfPower * x
	}
}
