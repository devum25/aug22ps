package darraygo

func TransposeSquare(A [][]int) [][]int {
	if len(A) == 0 {
		return nil
	}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			A[i][j], A[j][i] = A[j][i], A[i][j]
		}
	}

	return A
}
