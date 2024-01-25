package dfs

// visit all boundary elements and their  connected node and mark them as not to be acquired
// visit through all cell and convert all to be acquired to acquire and not to be acquired back to conflicted area
func AcquireConflictedArea(mat [][]byte) {

	for i := 0; i < len(mat); i++ { // left boundary
		if mat[i][0] == 'O' {
			markdfs(mat, i, 0)
		}
	}

	for i := 0; i < len(mat[0]); i++ { //top boundary
		if mat[0][i] == 'O' {
			markdfs(mat, 0, i)
		}
	}

	for j := 0; j < len(mat[0]); j++ { //bottom boundary
		if mat[len(mat)-1][j] == 'O' {
			markdfs(mat, len(mat)-1, j)
		}
	}

	for i := 0; i < len(mat); i++ { //right boundary
		if mat[i][len(mat[0])-1] == 'O' {
			markdfs(mat, i, len(mat[0])-1)
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 'O' {
				acquirekdfs(mat, i, j)
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 'K' {
				mat[i][j] = 'O'
			}
		}
	}
}

func markdfs(mat [][]byte, r, c int) {
	if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[r]) {
		return
	}
	row := []int{-1, 0, 1, 0}
	col := []int{0, -1, 0, -1}
	mat[r][c] = 'K'
	for i := 0; i < 4; i++ {
		if mat[r][c] == 'O' {
			markdfs(mat, r+row[i], c+col[i])
		}
	}
}

func acquirekdfs(mat [][]byte, r, c int) {
	if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[r]) {
		return
	}
	row := []int{-1, 0, 1, 0}
	col := []int{0, -1, 0, -1}
	mat[r][c] = 'X'
	for i := 0; i < 4; i++ {
		if mat[r][c] == 'O' {
			acquirekdfs(mat, r+row[i], c+col[i])
		}
	}
}
