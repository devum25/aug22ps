package backtracking

func SolveNQueensCount(N int) int {
	matrix := make([][]int, N)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, N)
	}
	hcol := make(map[int]bool)
	hrd := make(map[int]bool)
	hld := make(map[int]bool)
	res := 0
	placeNQueensPossibilitieeCount(&matrix, N, 0, &hcol, &hrd, &hld, &res)
	return res
}

func placeNQueensPossibilitieeCount(matrix *[][]int, N int, row int, hcol, hrd, hld *map[int]bool, res *int) bool {
	if row == N {
		*res++
		return true
	}

	for col := 0; col < N; col++ {
		if isValidBox(row, col, hcol, hrd, hld) {
			(*matrix)[row][col] = 'Q'
			(*hcol)[col] = true
			(*hrd)[col+row] = true
			(*hld)[col-row] = true
			placeNQueensPossibilitieeCount(matrix, N, row+1, hcol, hrd, hld, res)
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
func isValidBox(row, col int, hcol, hrd, hld *map[int]bool) bool {

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
