package arrayquestion

// Given an integer array nums and an integer k, return true if nums has a good subarray or false otherwise.

// A good subarray is a subarray where:

// its length is at least two, and
// the sum of the elements of the subarray is a multiple of k.
// Note that:

// A subarray is a contiguous part of the array.
// An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

// Example 1:

// Input: nums = [23,2,4,6,7], k = 6
// Output: true
// Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

func CheckSubArraySum(arr []int, k int) bool {
	pf := make([]int, len(arr))
	hash := make(map[int]int) // map of remainder and index
	hash[0] = -1
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			pf[i] = arr[0]
			if idx, ok := hash[pf[i]%k]; ok {
				if i-idx > 1 {
					return true
				}
			} else {
				hash[pf[i]%k] = i
			}
		} else {
			pf[i] = arr[i] + pf[i-1]

			if idx, ok := hash[pf[i]%k]; ok {
				if i-idx > 1 {
					return true
				}
			} else {
				hash[pf[i]%k] = i
			}
		}
	}

	return false
}
