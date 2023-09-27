package binarysearch

import (
	"math"
)

// Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.

// Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.

func NextGreaterElement(A int) int {

	res := make([]int, 0)
	a := A
	for a != 0 {
		res = append(res, a%10)
		a = a / 10
	}
	k := -1
	for i := 0; i < len(res); i++ {
		if i != 0 {
			if res[i] < res[i-1] {
				k = i
				break
			}
		}
	}
	num := math.MaxInt32
	j := -1
	for i := 0; i < k; i++ {
		if res[i] < num && res[i] > res[k] {
			num = res[i]
			j = i
		}
	}
	if k >= 0 && j >= 0 {
		res[k], res[j] = res[j], res[k]
	}

	for i := k - 1; i >= 0; i-- {
		min := res[i]
		k := -1
		for j := i; j >= 0; j-- {
			if res[j] < min {
				min = res[j]
				k = j
			}
		}
		if k >= 0 {
			res[i], res[k] = res[k], res[i]
		}

	}

	number := 0
	number = res[len(res)-1] * 10
	for i := len(res) - 2; i >= 0; i-- {
		number = (res[i] + number) * 10
	}
	if number/10 == A {
		return -1
	} else if number/10 > math.MaxInt32 {
		return -1
	}

	return number / 10

}
