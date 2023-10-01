package binarysearch

// Given an array of integers A and an integer B,
//  find and return the maximum value K such that there is no subarray in A of size K with the sum of elements greater than B.

func SpecialInteger(A []int, B int) int {
	l := 1
	r := len(A)
	ans := 0
	for l <= r {
		mid := (l + r) / 2

		if check(A, mid, B) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func check(A []int, k, b int) bool {

	sum := 0

	for i := 0; i < k; i++ {
		sum += A[i]
	}

	if sum > b {
		return false
	}

	for i := k; i < len(A); i++ {
		sum = sum + A[i]
		sum = sum - A[i-k]
		if sum > b {
			return false
		}
	}

	return true
}
