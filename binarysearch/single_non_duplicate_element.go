package binarysearch

// 540. Single Element in a Sorted Array
// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
// Return the single element that appears only once.
// Your solution must run in O(log n) time and O(1) space.

func SingleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if nums[0] != nums[1] {
		return nums[0]
	} else if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	l := 1
	r := len(nums) - 2

	for l <= r {
		mid := (l + r) / 2

		if nums[mid-1] != nums[mid] && nums[mid+1] != nums[mid] {
			return nums[mid]
		} else if mid%2 == 0 && nums[mid-1] == nums[mid] {
			r = mid - 2
		} else if mid%2 == 0 && nums[mid+1] == nums[mid] {
			l = mid + 2
		} else if mid%2 != 0 && nums[mid+1] == nums[mid] {
			r = mid - 1
		} else if mid%2 != 0 && nums[mid-1] == nums[mid] {
			l = mid + 1
		}
	}

	return -1
}
