package darraygo

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	processedMatrix := preprocess(matrix)
	return NumMatrix{Matrix: processedMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 && row2 == 0 && col2 == 0 {
		return this.Matrix[0][0]
	}
	if row1 == 0 && col1 == 0 {
		return this.Matrix[row2][col2]
	}
	// if row1 == row2 && col1 == col2 {
	// 	return this.Matrix[row1][col1]
	// }
	sum := 0
	if row1 == 0 {
		sum = this.Matrix[row2][col2] - this.Matrix[row2][col1-1]
	} else if col1 == 0 {
		sum = this.Matrix[row2][col2] - this.Matrix[row1-1][col2]
	} else {
		sum = this.Matrix[row2][col2] - this.Matrix[row1-1][col2] - this.Matrix[row2][col1-1] + this.Matrix[row1-1][col1-1]
	}

	return sum
}

func preprocess(matrix [][]int) [][]int {
	matrix = createrowwiseprefix(matrix)
	matrix = createcolwiseprefix(matrix)
	return matrix
}

func createcolwiseprefix(matrix [][]int) [][]int {
	for c := 0; c < len(matrix[0]); c++ {
		for r := 0; r < len(matrix); r++ {
			if r == 0 {
				continue
			}
			matrix[r][c] = matrix[r-1][c] + matrix[r][c]
		}
	}

	return matrix
}

func createrowwiseprefix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				continue
			}
			matrix[i][j] = matrix[i][j-1] + matrix[i][j]
		}
	}

	return matrix
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
