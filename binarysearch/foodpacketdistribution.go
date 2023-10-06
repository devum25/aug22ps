package binarysearch

import "math"

// The government wants to set up B distribution offices across N cities for the distribution of food packets.

// The population of the ith city is A[i]. Each city must have at least 1 office and every person is assigned to exactly one office in their own city.

// Let M denote the minimum number of people that are assigned to any of the offices. Find the maximum value of M possible.

func FoodPacketDistribution(A []int, B int) int {
	l := 1
	r := getMax(A)
	ans := math.MinInt
	for l <= r {
		mid := (l + r) / 2

		if checkOffice(A, mid, B) {

			ans = mid

			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func checkOffice(A []int, people, office int) bool {

	p := people
	k := 0
	for i := 0; i < len(A); i++ {
		if A[i]/p > 0 {
			k = k + (A[i] / p)
		}
	}

	if k >= office {
		return true
	}

	return false
}

func getMin(A []int) int {
	min := math.MaxInt

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
	}

	return min
}

func getMax(A []int) int {
	max := math.MinInt

	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}

	return max
}
