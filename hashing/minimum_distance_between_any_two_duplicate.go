package hashset

import "math"

func MinimumDistance(A []int) int {
	ans := math.MaxInt

	hash := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := hash[A[i]]; ok {
			dist := i - hash[A[i]]
			if dist < ans {
				ans = dist
			}
			hash[A[i]] = i
		} else {
			hash[A[i]] = i
		}
	}

	return ans
}
