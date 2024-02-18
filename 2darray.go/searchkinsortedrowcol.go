package darraygo

//Given a 2d matrix where row and cols are sorted in ascending order.
// Return true if k exist in 2d matrix.

func SearchK(A [][]int, k int) bool {

	r := 0
	c := 0

	for r < len(A) && c < len(A[0]) {
		if A[r][c] == k || A[r][len(A[0])-1] == k || A[len(A)-1][c] == k {
			return true
		} else if k > A[r][c] && k > A[len(A)-1][c] && k > A[r][len(A[0])-1] {
			r++
			c++
		} else {
			col := c
			row := r

			for row < len(A) {
				if A[row][col] == k {
					return true
				}
				row++
			}

			row = r
			for col < len(A[0]) {
				if A[row][col] == k {
					return true
				}
				col++
			}

			return false
		}
	}

	return false
}
