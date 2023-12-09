package backtracking

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// You may return the answer in any order.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
// Example 2:

// Input: n = 1
// Output: [["Q"]]

func NQueen(N int) bool {
	matrix := make([][]int, N)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, N)
	}
	hcol := make(map[int]bool)
	hrd := make(map[int]bool)
	hld := make(map[int]bool)
	return placeNQueens(&matrix, N, 0, &hcol, &hrd, &hld)
}

func placeNQueens(matrix *[][]int, N int, row int, hcol, hrd, hld *map[int]bool) bool {
	if row == N {
		return true
	}

	for col := 0; col < N; col++ {
		if isSafe(matrix, row, col, hcol, hrd, hld) {
			(*matrix)[row][col] = 1
			(*hcol)[col] = true
			(*hrd)[col+row] = true
			(*hld)[col-row] = true
			if placeNQueens(matrix, N, row+1, hcol, hrd, hld) {
				return true
			}
			(*matrix)[row][col] = 0
			delete(*hcol, col)
			delete(*hrd, col+row)
			delete(*hld, col-row)
		}
	}
	return false
}

// checks if already there is a queen in same row, same column, same diagonal
// return true is no queen is present in
// same row
// same column
// same diagonal
func isSafe(matrix *[][]int, row, col int, hcol, hrd, hld *map[int]bool) bool {

	if _, ok := (*hcol)[col]; ok {
		return false
	}

	if _, ok := (*hrd)[row+col]; ok {
		return false
	}

	if _, ok := (*hld)[col-row]; ok {
		return false
	}

	return true
}
