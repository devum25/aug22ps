package graphs

/**
 * @input A : 2D integer array
 * @input B : Integer array
 * @input C : Integer array
 *
 * @Output Integer
 */
//  Given a matrix of integers A of size N x M describing a maze. The maze consists of empty locations and walls.

//  1 represents a wall in a matrix and 0 represents an empty location in a wall.

//  There is a ball trapped in a maze. The ball can go through empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall (maze boundary is also considered as a wall). When the ball stops, it could choose the next direction.

//  Given two array of integers of size B and C of size 2 denoting the starting and destination position of the ball.

//  Find the shortest distance for the ball to stop at the destination. The distance is defined by the number of empty spaces traveled by the ball from the starting position (excluded) to the destination (included). If the ball cannot stop at the destination, return -1.

var dir [4][2]int = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

const inf = 1000000000

func inside(x, y, n, m int) bool {
	return (x >= 0 && x <= n-1 && y >= 0 && y <= m-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findMinimumDistance(maze [][]int, start, destination []int) int {
	n := len(maze)
	m := len(maze[0])
	dist := make([][][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = make([]int, 4)
			dist[i][j][0] = inf
			dist[i][j][1] = inf
			dist[i][j][2] = inf
			dist[i][j][3] = inf
		}
	}
	var q [][]int
	q = append(q, []int{start[0], start[1], 0})
	dist[start[0]][start[1]][0] = 0
	q = append(q, []int{start[0], start[1], 1})
	dist[start[0]][start[1]][1] = 0
	q = append(q, []int{start[0], start[1], 2})
	dist[start[0]][start[1]][2] = 0
	q = append(q, []int{start[0], start[1], 3})
	dist[start[0]][start[1]][3] = 0

	for len(q) != 0 {
		x, y, d := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		nx := x + dir[d][0]
		ny := y + dir[d][1]
		if inside(nx, ny, n, m) && maze[nx][ny] == 0 {
			if dist[nx][ny][d] == inf {
				dist[nx][ny][d] = dist[x][y][d] + 1
				q = append(q, []int{nx, ny, d})
			}
		} else {
			for idx, ad := range dir {
				nx := x + ad[0]
				ny := y + ad[1]
				if inside(nx, ny, n, m) && maze[nx][ny] == 0 && dist[nx][ny][idx] == inf {
					dist[nx][ny][idx] = dist[x][y][d] + 1
					q = append(q, []int{nx, ny, idx})
				}
			}
		}
	}
	ans := inf
	for i := 0; i < 4; i++ {
		x := destination[0]
		y := destination[1]
		if !inside(x+dir[i][0], y+dir[i][1], n, m) || maze[x+dir[i][0]][y+dir[i][1]] == 1 {
			ans = min(ans, dist[x][y][i])
		}
	}
	if ans == inf {
		ans = -1
	}
	return ans
}
