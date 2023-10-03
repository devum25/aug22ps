package binarysearch

import (
	"math"
)

// Given N tasks and k workers
// given an array of size N
// A[i] --> time to complete the ith task

// Find the minimum time to complete the task
// * One task can be done by 1 worker
// * A worker can only perform task which are contagious to each other.
// * All workers can do their work task in paraller

func MinimumTimeForKWorkersToCompleteNTask(N []int, k int) int {
	l, r := getMinMax(N)
	ans := math.MaxInt
	for l <= r {
		mid := (l + r) / 2

		if checkTime(N, mid, k) {
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

func checkTime(x []int, t, k int) bool {
	time := t
	worker := k
	spent := 0
	i := 0
	for ; i < len(x) && worker > 0; i++ {
		if spent+x[i] <= time {
			spent = spent + x[i]
			// lastRoom = A[i]
			// numcows++
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

func getMinMax(A []int) (int, int) {

	max := math.MinInt

	sum := 0

	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}

		sum += A[i]

	}
	return max, sum

}
