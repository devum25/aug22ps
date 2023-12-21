package backtracking

// Sudoku Solver
// Write a program to solve a Sudoku puzzle by filling the empty cells.

// A sudoku solution must satisfy all of the following rules:

// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
// The '.' character indicates empty cells.

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

// Output: [["5","3","4","6","7","8","9","1","2"],
// 		 ["6","7","2","1","9","5","3","4","8"],
// 		 ["1","9","8","3","4","2","5","6","7"],
//  		 ["8","5","9","7","6","1","4","2","3"],
// 		 ["4","2","6","8","5","3","7","9","1"],
// 		 ["7","1","3","9","2","4","8","5","6"],
// 		 ["9","6","1","5","3","7","2","8","4"],
// 		 ["2","8","7","4","1","9","6","3","5"],
// 		 ["3","4","5","2","8","6","1","7","9"]]
// Explanation: The input board is shown above and the only valid solution is shown below:

func solveSudoku(board [][]byte) {
	sudoku(&board, 0, 0, len(board))
}

func sudoku(matrix *[][]byte, i, j int, n int) bool {
	// base case
	if i == n {
		return true
	}
	// edge case
	if j == n {
		return sudoku(matrix, i+1, 0, n)
	}
	// if this cell is already populated
	if (*matrix)[i][j] != '.' {
		return sudoku(matrix, i, j+1, n)
	}

	for nu := 1; nu <= n; nu++ {
		if canbeplaced(matrix, i, j, byte(nu+'0')) {
			(*matrix)[i][j] = byte(nu + '0')
			if sudoku(matrix, i, j+1, n) {
				return true
			}
		}
	}

	(*matrix)[i][j] = '.'

	return false
}

func canbeplaced(matrix *[][]byte, i, j int, number byte) bool {
	// check for row and col
	for k := 0; k < len(*matrix); k++ {
		if (*matrix)[i][k] == number || (*matrix)[k][j] == number {
			return false
		}
	}
	// check for subgrid
	a := (i / 3) * 3
	b := (j / 3) * 3

	for x := a; x < a+3; x++ {
		for y := b; y < b+3; y++ {
			if (*matrix)[x][y] == number {
				return false
			}
		}
	}

	return true
}
