package backtracking

func SolveNQueens(n int) [][]string {
	board := make([][]rune, n)
	for r := 0; r < n; r++ {
		board[r] = make([]rune, n)
		for c := 0; c < n; c++ {
			board[r][c] = '.'
		}
	}
	hcol := make(map[int]bool)
	hrd := make(map[int]bool)
	hld := make(map[int]bool)
	res := [][]string{}
	placeNQueensPossibilities(&board, n, 0, &hcol, &hrd, &hld, &res)
	return res
}

func placeNQueensPossibilities(matrix *[][]rune, N int, row int, hcol, hrd, hld *map[int]bool, res *[][]string) bool {
	if row == N {
		appendResult(matrix, res)
		return true
	}

	for col := 0; col < N; col++ {
		if isValid(row, col, hcol, hrd, hld) {
			(*matrix)[row][col] = 'Q'
			(*hcol)[col] = true
			(*hrd)[col+row] = true
			(*hld)[col-row] = true
			placeNQueensPossibilities(matrix, N, row+1, hcol, hrd, hld, res)
			// return true

			(*matrix)[row][col] = '.'
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
func isValid(row, col int, hcol, hrd, hld *map[int]bool) bool {

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

func appendResult(matrix *[][]rune, res *[][]string) {
	var temp []string
	for i := 0; i < len(*matrix); i++ {
		var s string
		for j := 0; j < len(*matrix); j++ {
			s += string((*matrix)[i][j])
		}
		temp = append(temp, s)
	}
	*res = append(*res, temp)
}
