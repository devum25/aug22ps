package graphs

import "math"

// 994. Rotting Oranges
// Medium
// Topics
// Companies
// You are given an m x n grid where each cell can have one of three values:

// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

type Oranges struct {
	Crd      Coordinate
	Distance int
}

func OrangesRotting(grid [][]int) int {
	queue := make([]Oranges, 0)
	ans := math.MinInt
	hash := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		hash[i] = make([]int, len(grid[i]))
	}

	res := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		res[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res[i][j] = -1
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				hash[i][j] = 1
				queue = append(queue, Oranges{Crd: Coordinate{x: i, y: j}, Distance: 0})
			}
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		i := temp.Crd.x
		j := temp.Crd.y

		if (i-1) >= 0 && hash[i-1][j] == 0 && grid[i-1][j] == 1 {
			queue = append(queue, Oranges{Crd: Coordinate{x: i - 1, y: j}, Distance: temp.Distance + 1})
			hash[i-1][j] = 1
			res[i-1][j] = temp.Distance + 1
		}
		if (j-1) >= 0 && hash[i][j-1] == 0 && grid[i][j-1] == 1 {
			queue = append(queue, Oranges{Crd: Coordinate{x: i, y: j - 1}, Distance: temp.Distance + 1})
			hash[i][j-1] = 1
			res[i][j-1] = temp.Distance + 1
		}
		if (i+1) < len(grid) && hash[i+1][j] == 0 && grid[i+1][j] == 1 {
			queue = append(queue, Oranges{Crd: Coordinate{x: i + 1, y: j}, Distance: temp.Distance + 1})
			hash[i+1][j] = 1
			res[i+1][j] = temp.Distance + 1
		}
		if (j+1) < len(grid[i]) && hash[i][j+1] == 0 && grid[i][j+1] == 1 {
			queue = append(queue, Oranges{Crd: Coordinate{x: i, y: j + 1}, Distance: temp.Distance + 1})
			hash[i][j+1] = 1
			res[i][j+1] = temp.Distance + 1
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if res[i][j] == -1 && grid[i][j] == 1 {
				return -1
			} else {
				ans = max(ans, res[i][j])
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
