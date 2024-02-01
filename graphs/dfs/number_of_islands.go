package dfs

// traverse graph using dfs, can also be done using bfs
// mark visted nodes as visted by changing the value inside matrix as 0
func NumberOfIsland(mat [][]int) int {
	islands := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				dfsWithDiagonal(mat, i, j)
				islands++
			}
		}
	}

	return islands
}

func dfs2(mat [][]int, r, c int) {
	if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[r]) || mat[r][c] == 0 {
		return
	}
	mat[r][c] = 0
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		dfs2(mat, i+row[i], c+col[i])
	}
}

func dfsWithDiagonal(mat [][]int, r, c int) {
	if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[r]) || mat[r][c] == 0 {
		return
	}

	mat[r][c] = 0

	row := []int{-1, 0, 1, 0, -1, 1, 1, -1}
	col := []int{0, 1, 0, -1, 1, 1, -1, -1}

	for i := 0; i < 8; i++ {
		dfsWithDiagonal(mat, r+row[i], c+col[i])
	}
}
