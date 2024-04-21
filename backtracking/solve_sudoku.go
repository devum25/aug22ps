package backtracking

import "strconv"

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

func SolveSudoku(board [][]byte) {

	grid := transform(board)

	solve3(grid, 0, 0)

	transform1(grid, board)
}

func solve3(board [][]int, row, col int) bool {
	if col == len(board) {
		col = 0
		row++
	}

	if row == len(board) {
		return true
	}

	if board[row][col] > 0 {
		return solve3(board, row, col+1)
	}

	for i := 1; i <= 9; i++ {
		if isValid1(board, row, col, i) {
			board[row][col] = i
			if solve3(board, row, col+1) {
				return true
			}
			board[row][col] = 0
		}
	}

	return false
}

func isValid1(grid [][]int, row, col, val int) bool {
	for k := 0; k < len(grid); k++ {
		if grid[row][k] == val || grid[col][k] == val {
			return false
		}
	}

	u := (row / 3) * 3
	v := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[u+i][v+j] == val {
				return false
			}
		}
	}

	return true
}

func transform1(grid [][]int, board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = byte(grid[i][j] + '0')
		}
	}
}

func transform(board [][]byte) [][]int {
	grid := make([][]int, len(board))

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			grid[i][j], _ = strconv.Atoi(string(board[i][j]))
		}
	}

	return grid
}
