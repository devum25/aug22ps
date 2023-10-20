package recursion

func Solve(a, b int) int {
	arr := make([]int, 0)
	for i := 0; i < a; i++ {
		arr = append(arr, i+1)
	}

	return helpJoseph(0, b, b-1, len(arr), arr)
}

func helpJoseph(idx int, k int, enemy int, n int, cir []int) int {
	if len(cir) == 1 {
		return cir[0]
	}
	cir1 := make([]int, 0)
	cir1 = append(cir1, cir[:enemy]...)
	if enemy+1 < n {
		cir1 = append(cir1, cir[enemy+1:n]...)
	}
	n = len(cir1)
	idx = enemy % n
	enemy = ((enemy % n) + (k - 1)) % n
	x := helpJoseph(idx, k, enemy, len(cir1), cir1)
	return x
}

// better way
func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}

	m := []int{}
	// 0 based indexing
	for i := 0; i < n; i++ {
		m = append(m, i)
	}

	start := 0
	for len(m) != 1 {
		out := (start + k - 1) % len(m)
		// delete element at index k
		m = append(m[:out], m[out+1:]...)
		start = out
	}

	return m[0] + 1
}
