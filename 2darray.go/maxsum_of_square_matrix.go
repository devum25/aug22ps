package darraygo

import "math"

// Given a 2D integer matrix A of size N x N, find a B x B submatrix where B<= N and B>= 1,
//  such that the sum of all the elements in the submatrix is maximum.

func MaxSubSquareMatrix(A [][]int, B int) int {

	preProcess := make([][]int, len(A)-B+1)

	for i := 0; i < len(preProcess); i++ {
		preProcess[i] = make([]int, len(A[0]))
	}

	for col := 0; col < len(A[0]); col++ {
		sum := 0
		for row := 0; row < B; row++ {
			sum += A[row][col]
		}
		preProcess[0][col] = sum
		for row := 1; row < B; row++ {
			preProcess[row][col] = sum - A[row-1][col] + A[row+B-1][col]
			sum = preProcess[row][col]
		}
	}
	max := math.MinInt
	for i := 0; i < len(preProcess); i++ {
		sum := 0
		for j := 0; j < B; j++ {
			sum = sum + preProcess[i][j]
		}

		if sum > max {
			max = sum
		}

		for j := 1; j < len(preProcess[i])-B+1; j++ {
			sum = sum - preProcess[i][j-1] + preProcess[i][j+B-1]
			if sum > max {
				max = sum
			}
		}

	}

	return max

}
