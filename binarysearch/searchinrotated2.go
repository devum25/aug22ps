package binarysearch

import "math"

// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

// You must decrease the overall operation steps as much as possible.

func SearchWithDuplicate(A []int, target int) bool {

	pivot := getMaxIdx(A)

	if A[0] >= target {
		l := pivot + 1
		r := len(A) - 1

		for l <= r {
			mid := (l + r) / 2

			if A[mid] == target {
				return true
			} else if A[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	if A[0] <= target {
		l := 0
		r := pivot
		for l <= r {
			mid := (l + r) / 2
			if A[mid] == target {
				return true
			} else if A[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

func getMaxIdx(A []int) int {
	max := math.MinInt
	idx := -1
	for i := 0; i < len(A); i++ {
		if A[i] >= max && (((i+1) < len(A) && A[i] > A[i+1]) || (i == (len(A) - 1))) {
			max = A[i]
			idx = i
		}
	}

	return idx
}
