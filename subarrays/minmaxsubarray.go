package subarrays

import "math"

func MinLenghtSubArrayHavingMinMax(A []int) int {
	maxi := -1
	mini := -1

	max := math.MinInt

	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}

	min := math.MaxInt

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
	}

	if min == max {
		return 1
	}

	res := math.MaxInt

	for i := len(A) - 1; i >= 0; i-- {

		if A[i] == max {

			if mini == -1 {
				maxi = i
				continue
			}
			maxi = i

			if res > (int(math.Abs(float64(mini-maxi))))+1 {
				res = int(math.Abs(float64(mini-maxi))) + 1
			}

		} else if A[i] == min {

			if maxi == -1 {
				mini = i
				continue
			}

			mini = i
			if res > (int(math.Abs(float64(mini-maxi))) + 1) {
				res = int(math.Abs(float64(mini-maxi))) + 1
			}
		}

	}

	return res
}

func getMaxVal(arr []int) int {
	max := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
func getMinVal(arr []int) int {
	min := math.MaxInt

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}
