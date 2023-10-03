package binarysearch

import "math"

/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer array
 *
 * @Output Integer
 */
func Paint(A int, B int, C []int) int {
	l, r := getMinMax(C)
	ans := math.MaxInt
	for l <= r {
		mid := (l + r) / 2

		if checkPainter(C, mid, A, B) {
			if mid < ans {
				ans = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return (ans * B) % 10000003
}

func checkPainter(x []int, t, k, unit int) bool {
	time := t
	worker := k
	spent := 0
	i := 0
	for ; i < len(x) && worker > 0; i++ {
		if spent+x[i] <= time {
			spent = spent + x[i]
			if i == len(x)-1 {
				return true
			}
		} else {
			worker--
			spent = 0
			i--
		}

	}

	return false
}

func getMinPainter(A []int) int {
	min := math.MaxInt

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}

	}
	return min

}

func getMaxPainter(A []int) int {
	min := math.MinInt

	for i := 0; i < len(A); i++ {
		if A[i] > min {
			min = A[i]
		}

	}
	return min

}
