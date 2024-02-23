package arrayquestion

import (
	"math"
	"sort"
)

// Given an array A of N integers where the i-th element represent the number of chocolates in the i-th packet.

// There are B number of students, the task is to distribute chocolate packets following below conditions:

// 1. Each student gets one packets.
// 2. The difference between the number of chocolates given to any two students is minimum.
// Return the minimum difference (that can be achieved) between the student who gets minimum number of chocolates and the student who gets maximum number of chocolates.
func Distribute(A []int, B int) int {
	sort.Ints(A)

	n := B - 1

	min := math.MaxInt

	min = A[n] - A[0]
	i := 1
	j := n + 1
	for j < len(A) {
		x := A[j] - A[i]

		if x < min {
			min = x
		}
		j++
		i++
	}

	return min
}
