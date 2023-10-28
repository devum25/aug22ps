package subarrays

func IsEquillibrium(A []int) int {

	pf := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		if i == 0 {
			pf[i] = A[i]
		} else {
			pf[i] = A[i] + pf[i-1]
		}
	}

	for i := 0; i < len(A); i++ {
		left := 0
		right := 0

		left = getSum(pf, 0, i-1)
		right = getSum(pf, i+1, len(A)-1)

		if left == right {
			return i
		}
	}

	return -1

}

func getSum(pf []int, s, e int) int {
	if s == 0 && e < 0 {
		return 0
	} else if s >= len(pf) {
		return 0
	} else if s == 0 {
		return pf[e]
	} else {
		return pf[e] - pf[s-1]
	}
}
