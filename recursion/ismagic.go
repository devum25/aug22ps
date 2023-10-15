package recursion

func IsMagic(x int) int {

	y := x

	for y/10 != 0 {
		y = ismagic(y)
	}

	if y == 1 {
		return 1
	}

	return 0
}

func ismagic(x int) int {
	if x == 0 {
		return 0
	}

	ans := x%10 + ismagic(x/10)
	return ans
}
