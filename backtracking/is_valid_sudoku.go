package backtracking

// 36. Valid Sudoku
// Medium
// 10K
// 1K
// Companies
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Example 1:
// Input: board = [["5","3",".",".","7",".",".",".","."],
// 				["6",".",".","1","9","5",".",".","."],
// 				[".","9","8",".",".",".",".","6","."],
// 				["8",".",".",".","6",".",".",".","3"],
// 				["4",".",".","8",".","3",".",".","1"],
// 				["7",".",".",".","2",".",".",".","6"],
// 				[".","6",".",".",".",".","2","8","."],
// 				[".",".",".","4","1","9",".",".","5"],
// 				[".",".",".",".","8",".",".","7","9"]]

func ValidateSudoku(board [][]byte) bool {
	valid := true
	validateSudoku(&board, 0, 0, len(board), &valid)
	return valid
}

func validateSudoku(matrix *[][]byte, i, j, n int, valid *bool) {
	if i == n {
		*valid = true
		return
	}

	if j == n {
		validateSudoku(matrix, i+1, 0, n, valid)
		return
	}

	if (*matrix)[i][j] == '.' {
		validateSudoku(matrix, i, j+1, n, valid)
		return
	}

	if int((*matrix)[i][j]-'0') > 9 {
		*valid = false
		return
	}

	if *valid && verify(matrix, i, j, (*matrix)[i][j]) {
		validateSudoku(matrix, i, j+1, n, valid)
		return
	} else {
		*valid = false
	}

	return
}

func verify(matrix *[][]byte, i, j int, number byte) bool {
	// check for row and col
	for row := i + 1; row < len(*matrix); row++ {
		if (*matrix)[row][j] == number {
			return false
		}
	}

	for row := 0; row < i-1; row++ {
		if (*matrix)[row][j] == number {
			return false
		}
	}

	for col := j + 1; col < len(*matrix); col++ {
		if (*matrix)[i][col] == number {
			return false
		}
	}

	for col := 0; col < j-1; col++ {
		if (*matrix)[i][col] == number {
			return false
		}
	}
	// check for subgrid
	a := (i / 3) * 3
	b := (j / 3) * 3

	for x := a; x < a+3; x++ {
		for y := b; y < b+3; y++ {
			if i == x && j == y {
				continue
			}
			if (*matrix)[x][y] == number {
				return false
			}
		}
	}

	return true
}
