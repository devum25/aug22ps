package arrayquestion

import "math"

// 53. Maximum Subarray

// Given an integer array nums, find the
// subarray
//  with the largest sum, and return its sum.

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6

// The subarray with negative sum is discarded (by assigning max_ending_here = 0 in code).
// We carry subarray till it gives positive sum.

func MaxSumSubArray(arr []int) int {
	ans := math.MinInt

	// s := 0
	e := 0

	max_sum := 0

	for e < len(arr) {
		max_sum = max_sum + arr[e]
		if max_sum > ans {
			ans = max_sum
		}
		if max_sum < 0 {
			max_sum = 0
		}

		e++
	}

	return ans
}
