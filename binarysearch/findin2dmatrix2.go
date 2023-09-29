package binarysearch

// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.
func FindIn2DMatrix(A [][]int, x int) bool {

	// rows := make([]int, 0)
	for i := len(A[0]) - 1; i >= 0; i-- {
		if A[0][i] > x {
			continue
		} else if A[0][i] == x {
			return true
		} else {
			found := searchhere(A, 0, len(A)-1, i, x)
			if found {
				return found
			}
			continue
		}
	}

	return false
}

func searchhere(A [][]int, l, r, col, x int) bool {

	for l <= r {
		mid := (l + r) / 2

		if A[mid][col] == x {
			return true
		} else if A[mid][col] > x {
			r = mid - 1
		} else if A[mid][col] < x {
			l = mid + 1
		}
	}

	return false
}
