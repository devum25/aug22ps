package slidingwindow

func Sumofsubarrayofsizek() []int {
	A := []int{-3, 4, -2, 5, 3, -2, 8, 2, -1, 4}
	res := make([]int, 0)
	k := 4
	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + A[i]
	}
	res = append(res, sum)
	s := 1
	e := k

	for e < len(A) {
		sum = sum - A[s-1] + A[e]
		res = append(res, sum)

		s++
		e++
	}

	return res
}
