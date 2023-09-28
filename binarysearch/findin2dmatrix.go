package binarysearch

import "fmt"

// matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
func SearchMatrix(A [][]int, x int) bool {
	// col := len(A[])
	col := len(A[0]) - 1
	resCol := make([]int, 0)
	for i := 0; i < len(A); i++ {
		resCol = append(resCol, A[i][0])
	}

	l := 0
	r := len(resCol) - 1
	fmt.Print(col)
	for l <= r {
		mid := (l + r) / 2
		if A[mid][0] == x || A[mid][col] == x {
			return true
		} else if A[mid][0] < x && x < A[mid][col] {
			return searchinthisrow(A[mid], 0, col, x)
		} else if x < A[mid][0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return true
}

func searchinthisrow(A []int, l, r, x int) bool {

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == x {
			return true
		} else if A[mid] < x {
			l = mid + 1
		} else if A[mid] > x {
			r = mid - 1
		}
	}

	return false
}
