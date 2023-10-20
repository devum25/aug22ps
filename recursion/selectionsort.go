package recursion

import "math"

func SelectionSort(n []int) []int {
	x := []int{}
	helpersort(&n, &x)
	return x
}

func helpersort(n *[]int, x *[]int) {
	y := *n
	if len(y) == 1 {
		*x = append(*x, y[0])
		return
	}

	*x = append(*x, min(n))

	helpersort(n, x)
}

func min(n *[]int) int {
	min := math.MaxInt
	idx := 0
	y := *n
	for i := 0; i < len(y); i++ {
		if y[i] < min {
			min = y[i]
			idx = i
		}
	}

	y[0], y[idx] = y[idx], y[0]

	y = y[1:]
	*n = y
	return min
}
