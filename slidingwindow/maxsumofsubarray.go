package slidingwindow

import "math"

func MaxSumOfSubArray() int {
	A := []int{-3, 4, -2, 5, 3, -2, 8, 2, -1, 4}
	k := 4

	sum := math.MinInt
	x := 0
	for i := 0; i < k; i++ {
		x += A[i]
	}

	if x > sum {
		sum = x
	}

	s := 1
	e := k

	for e < len(A) {
		x = x - A[s-1] + A[e]
		s++
		e++
		if x > sum {
			sum = x
		}
	}

	return sum
}
