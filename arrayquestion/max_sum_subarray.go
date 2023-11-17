package arrayquestion

import "math"

func MaxSubArray(arr []int) int {
	pf := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			pf[i] = arr[i]
		} else {
			pf[i] = arr[i] + pf[i-1]
		}
	}

	max := math.MinInt

	for i := 0; i < len(arr); i++ {
		tmp := 0
		for j := i; j < len(arr); j++ {
			if i == 0 {
				tmp = pf[j]
			} else {
				tmp = pf[j] - pf[i-1]
			}

			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}
