package graphs

// Given a matrix of integers A of size N x M consisting of 0 or 1.
// For each cell of the matrix find the distance of nearest 1 in the matrix.
// Distance between two cells (x1, y1) and (x2, y2) is defined as |x1 - x2| + |y1 - y2|.
// Find and return a matrix B of size N x M which defines for each cell in A distance of nearest 1 in the matrix A.
// NOTE: There is atleast one 1 is present in the matrix

func NearestCell(mat [][]int) [][]int {
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
			if mat[i][j] == 1 {
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

		if (i-1) >= 0 && hash[i-1][j] == 0 {
			queue = append(queue, Building{Crd: Coordinate{x: i - 1, y: j}, Distance: temp.Distance + 1})
			hash[i-1][j] = 1
			res[i-1][j] = temp.Distance + 1
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
