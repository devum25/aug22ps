package darraygo

// Given a 2d matrix, calculate sum of all possible submatrix.

// Can be solved using contribution technique, find the number of submatrix where element (i,j) will fall.

// count number of TopLeft * count number of BottomRight * A[i][j]

func TotalSumSubmatrix(A [][]int) int {
	sum := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			sum += A[i][j] * (i + 1) * (j + 1) * (len(A) - i) * (len(A[i]) - j)
		}
	}

	return sum
}
