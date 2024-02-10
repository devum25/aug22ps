package graphs

// Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

// A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

// All the visited cells of the path are 0.
// All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
// The length of a clear path is the number of visited cells of this path.

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	} else if grid[0][0] == 0 && len(grid) == 1 && len(grid[0]) == 1 {
		return 1
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				return bfs1(grid, i, j)
			}
		}
	}

	return -1

}

func bfs1(grid [][]int, r, c int) int {
	row := []int{-1, 0, 1, 0, -1, 1, -1, 1}
	col := []int{0, 1, 0, -1, -1, 1, 1, -1}
	grid[r][c] = 1
	queue := make([]Node1, 0)
	queue = append(queue, Node1{Distance: 1, Crd: Coordinate1{x: r, y: c}})
	mat := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		mat[i] = make([]int, len(grid[i]))
	}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		dist := top.Distance
		i := top.Crd.x
		j := top.Crd.y
		for k := 0; k < 8; k++ {
			if (i+row[k]) < 0 || (i+row[k]) >= len(grid) || (j+col[k]) < 0 || (j+col[k]) >= len(grid[i+row[k]]) || grid[i+row[k]][j+col[k]] == 1 {
				continue
			}
			grid[i+row[k]][j+col[k]] = 1
			queue = append(queue, Node1{Distance: dist + 1, Crd: Coordinate1{x: i + row[k], y: j + col[k]}})
			mat[i+row[k]][j+col[k]] = dist + 1
		}

	}

	if mat[len(mat)-1][len(mat[0])-1] == 0 {
		return -1
	}

	return mat[len(mat)-1][len(mat[0])-1]

}

type Coordinate1 struct {
	x int
	y int
}

type Node1 struct {
	Crd      Coordinate1
	Distance int
}
