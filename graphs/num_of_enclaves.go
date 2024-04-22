package graphs

func NumEnclaves(grid [][]int) int {
	total := 0

	queue := make([]Coordinate, 0)
	visited := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				total++
			}
		}
	}

	row := 0
	col := 0

	for col < len(grid[row]) {
		if grid[row][col] == 1 && visited[row][col] != -1 {
			c := Coordinate{x: row, y: col}
			queue = append(queue, c)
			visited[row][col] = -1
			total--
		}
		col++
	}

	col = 0
	for row < len(grid) {
		if grid[row][col] == 1 && visited[row][col] != -1 {
			c := Coordinate{x: row, y: col}
			queue = append(queue, c)
			visited[row][col] = -1
			total--
		}
		row++
	}

	row = len(grid) - 1
	col = 0
	for col < len(grid[0]) {
		if grid[row][col] == 1 && visited[row][col] != -1 {
			c := Coordinate{x: row, y: col}
			queue = append(queue, c)
			visited[row][col] = -1
			total--
		}
		col++
	}

	col = len(grid[0]) - 1

	for row >= 0 {
		if grid[row][col] == 1 && visited[row][col] != -1 {
			c := Coordinate{x: row, y: col}
			queue = append(queue, c)
			visited[row][col] = -1
			total--
		}
		row--
	}

	rows := []int{-1, 0, 1, 0}
	cols := []int{0, -1, 0, 1}
	ans := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		r := node.x
		c := node.y

		for k := 0; k < 4; k++ {
			if (r+rows[k]) < 0 || (r+rows[k]) >= len(grid) || (c+cols[k]) < 0 || (c+cols[k]) >= len(grid[0]) || visited[r+rows[k]][c+cols[k]] == -1 || grid[r+rows[k]][c+cols[k]] != 1 {
				continue
			}

			visited[r+rows[k]][c+cols[k]] = -1
			queue = append(queue, Coordinate{x: r + rows[k], y: c + cols[k]})
			ans++
		}
	}

	return total - ans
}
