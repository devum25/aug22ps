package arrayquestion

import "math"

// [2,-5,1,-4,3,-2] --> [-2,5,-1,4,-3,2] [-5,1,-4]

func MaxAbsoluteSum(arr []int) int {
	max := getMax(arr)
	min := getMin(arr)

	ans := math.Max(float64(min), float64(max))

	return int(ans)
}

func getMax(arr []int) int {
	max := math.MinInt

	e := 0
	curr := 0

	for e < len(arr) {
		curr = curr + arr[e]
		if curr > max {
			max = curr
		}

		if curr < 0 {
			curr = 0
		}
		e++
	}

	return max
}

func getMin(arr []int) int {
	min := math.MaxInt

	e := 0

	curr := 0

	for e < len(arr) {
		curr = curr + arr[e]
		if curr < min {
			min = curr
		}

		if curr > 0 {
			curr = 0
		}
		e++
	}

	return int(math.Abs(float64(min)))
}
