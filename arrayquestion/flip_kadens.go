package arrayquestion

import (
	"math"
	"strconv"
)

// Problem Description
// You are given a binary string A(i.e., with characters 0 and 1) consisting of characters A1, A2, ..., AN. In a single operation, you can choose two indices, L and R, such that 1 ≤ L ≤ R ≤ N and flip the characters AL, AL+1, ..., AR. By flipping, we mean changing character 0 to 1 and vice-versa.

// Your aim is to perform ATMOST one operation such that in the final string number of 1s is maximized.

// If you don't want to perform the operation, return an empty array. Else, return an array consisting of two elements denoting L and R. If there are multiple solutions, return the lexicographically smallest pair of L and R.

// NOTE: Pair (a, b) is lexicographically smaller than pair (c, d) if a < c or, if a == c and b < d.

func Flip(a string) []int {
	arr := convertStringToArr(a)

	max := math.MinInt

	e := 0
	s := 0

	res := make([]int, 2)
	curr := 0
	for e < len(arr) {
		if arr[e] == 1 {
			curr = -1 + curr
			if curr < 0 {
				s = e + 1
				curr = 0
			}
		} else {
			curr = curr + arr[e] + 1
			if curr > max {
				max = curr
				res[0] = s + 1
				res[1] = e + 1
			} else if curr == max {
				if (s + 1) < res[0] {
					res[0] = s + 1
					res[1] = e + 1
				}
			}
		}
		e++
	}

	if res[0] == res[1] && res[1] == 0 {
		return nil
	}

	return res
}

func convertStringToArr(a string) []int {
	res := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(string(a[i]))
		res[i] = val
	}

	return res
}
