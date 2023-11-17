package arrayquestion

func MaxProductSubarray(arr []int) int {
	max := arr[0]
	min := arr[0]
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		temp := max

		max = maximum(arr[i], max*arr[i], min*arr[i])
		min = minimum(arr[i], temp*arr[i], min*arr[i])

		if max > ans {
			ans = max
		}

		if arr[i] == 0 {
			max = 0
			min = 0
		}
	}

	return ans
}

func maximum(a, b, c int) int {
	if (a > b && a >= c) || (a > c && a >= b) {
		return a
	} else if (b > a && b >= c) || (b >= a && b > c) {
		return b
	} else {
		return c
	}
}

func minimum(a, b, c int) int {
	if (a <= b && a < c) || (a <= c && a < b) {
		return a
	} else if (b < a && b <= c) || (b < c && b <= a) {
		return b
	} else {
		return c
	}
}
