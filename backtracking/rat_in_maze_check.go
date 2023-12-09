package backtracking

//Given the maze,return true is there exists a path from the rat to cheeze.
// 0 in maze represents -- non-blocked
// 1 in maze represents -- blocked

func MazeSolver_CheckPath(A [][]int) bool {
	return false
}

func mazesolver(matrix [][]int, x, y, N, M int) bool {
	if x == N-1 && y == M-1 {
		return true
	}
	if x >= N || y >= M || x < 0 || y < 0 {
		return false
	}
	if matrix[x][y] == 1 || matrix[x][y] == 2 {
		return false
	}

	matrix[x][y] = 2
	return mazesolver(matrix, x+1, y, N, M) || mazesolver(matrix, x, y+1, N, M) || mazesolver(matrix, x-1, y, N, M) || mazesolver(matrix, x, y-1, N, M)

}
