package graphs

func OrangesRotting1(grid [][]int) int {
	queue := make([]Oranges, 0)
	visited := make([][]int, len(grid))
	count := 0
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				c := Coordinate{x: i, y: j}
				visited[i][j] = -1
				queue = append(queue, Oranges{Crd: c})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	row := []int{-1, 0, 1, 0}
	col := []int{0, -1, 0, 1}

	ans := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		r := node.Crd.x
		c := node.Crd.y
		d := node.Distance

		for k := 0; k < 4; k++ {
			if (r+row[k]) < 0 || (c+col[k]) < 0 || (r+row[k]) >= len(grid) || (c+col[k]) >= len(grid[0]) || visited[r+row[k]][c+col[k]] == -1 || grid[r+row[k]][c+col[k]] != 1 {
				continue
			}
			visited[r+row[k]][c+col[k]] = -1
			c := Coordinate{x: r + row[k], y: c + col[k]}
			queue = append(queue, Oranges{Crd: c, Distance: 1 + d})
			ans = (1 + d)
			count--
		}
	}

	if count > 0 {
		return -1
	}

	return ans
}

// type Coordinate struct{
//     x int
//     y int
// }

// type Oranges struct{
//     C Coordinate
//     Dist int
// }
