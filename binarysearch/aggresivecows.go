package binarysearch

import "sort"

func AggresiveCows(A []int, B int) int {
	sort.Ints(A)
	l := 0
	r := A[len(A)-1] - A[0]
	ans := 0
	for l <= r {

		mid := (l + r) / 2

		if Check(A, mid, B) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func Check(A []int, d, n int) bool {
	lastRoom := A[0]
	numcows := 1
	for i := 1; i < len(A); i++ {
		if A[i]-lastRoom >= d {
			lastRoom = A[i]
			numcows++
		}

		if n == numcows {
			return true
		}
	}

	return false
}
