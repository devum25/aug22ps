package darraygo

func SpiralOfRecatange(A [][]int) []int {
	res := make([]int, 0)
	row := 0
	col := 0
	R := len(A)
	C := len(A[0])

	for C > 1 && R > 1 {

		for k := 1; k < C; k++ {
			res = append(res, A[row][col])
			col++
		}

		for k := 1; k < R; k++ {
			res = append(res, A[row][col])
			row++
		}

		for k := 1; k < C; k++ {
			res = append(res, A[row][col])
			col--
		}

		for k := 1; k < R; k++ {
			res = append(res, A[row][col])
			row--
		}

		C = C - 2
		R = R - 2
		row++
		col++
	}

	if C == 1 && R > 1 {
		i := row
		for R > 0 {
			res = append(res, A[i][col])
			i++
			R--
		}
	} else if R == 1 && C > 1 {
		i := col
		for C > 0 {
			res = append(res, A[row][i])
			i++
			C--
		}
	} else if R == 1 && C == 1 {
		res = append(res, A[row][col])
	}

	return res
}
