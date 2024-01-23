package graphs

// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.

// Similar to house hospital problem, consider 0 as hospital and 1 as house

type Building struct {
	Crd      Coordinate
	Distance int
}

type Coordinate struct {
	x int
	y int
}

func UpdateMatrix(mat [][]int) [][]int {
	queue := make([]Building, 0)
	hash := make([][]int, len(mat))
	res := make([][]int, len(mat))
	for i := 0; i < len(hash); i++ {
		hash[i] = make([]int, len(mat[i]))
	}
	for i := 0; i < len(hash); i++ {
		res[i] = make([]int, len(mat[i]))
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				Crd := Coordinate{i, j}
				hash[i][j] = 1
				queue = append(queue, Building{Crd: Crd, Distance: 0})
			}
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		i := temp.Crd.x
		j := temp.Crd.y

		if (i - 1) >= 0 {
			if mat[i-1][j] != 0 && hash[i-1][j] == 0 {
				queue = append(queue, Building{Crd: Coordinate{x: i - 1, y: j}, Distance: temp.Distance + 1})
				hash[i-1][j] = 1
				res[i-1][j] = temp.Distance + 1
			}
		}
		if (j-1) >= 0 && hash[i][j-1] == 0 {
			queue = append(queue, Building{Crd: Coordinate{x: i, y: j - 1}, Distance: temp.Distance + 1})
			hash[i][j-1] = 1
			res[i][j-1] = temp.Distance + 1
		}
		if (i+1) < len(mat) && hash[i+1][j] == 0 {
			queue = append(queue, Building{Crd: Coordinate{x: i + 1, y: j}, Distance: temp.Distance + 1})
			hash[i+1][j] = 1
			res[i+1][j] = temp.Distance + 1
		}
		if (j+1) < len(mat[i]) && hash[i][j+1] == 0 {
			queue = append(queue, Building{Crd: Coordinate{x: i, y: j + 1}, Distance: temp.Distance + 1})
			hash[i][j+1] = 1
			res[i][j+1] = temp.Distance + 1
		}
	}

	return res
}
