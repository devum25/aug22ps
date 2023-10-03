package binarysearch

import "math"

func AllocateBooks(A []int, B int) int {

	if len(A) < B {
		return -1
	}

	l, r := getMinMax(A)
	ans := math.MaxInt
	for l <= r {
		mid := (l + r) / 2

		if checkDistribution(A, B, mid) {
			if mid < ans {
				ans = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func checkDistribution(A []int, s, p int) bool {
	students := s
	pages := p
	distributed := 0
	for i := 0; i < len(A) && students > 0; i++ {

		if distributed+A[i] <= pages && ((len(A)-1)-i) >= (students-1) {
			distributed = distributed + A[i]

			if i == len(A)-1 {
				return true
			}
		} else {
			students--
			distributed = 0
			i--
		}
	}

	return false
}
