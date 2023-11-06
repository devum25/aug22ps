package darraygo

func Spiral(A [][]int) {
	res := make([]int, 0)
	row := 0
	col := 0
	N := len(A)
	for N > 1 {

		for k := 1; k < N; k++ {
			res = append(res, A[row][col])
			col++
		}

		for k := 1; k < N; k++ {
			res = append(res, A[row][col])
			row++
		}

		for k := 1; k < N; k++ {
			res = append(res, A[row][col])
			col--
		}

		for k := 1; k < N; k++ {
			res = append(res, A[row][col])
			row--
		}

		N = N - 2
		row++
		col++
	}

	if N == 1 {
		res = append(res, A[row][col])
	}
}
