package arrayquestion

import "math"

func AddedInteger(nums1 []int, nums2 []int) int {
	minNums1 := math.MaxInt

	minNums2 := math.MaxInt

	for i := 0; i < len(nums1); i++ {
		if nums1[i] < minNums1 {
			minNums1 = nums1[i]
		}
		if nums2[i] < minNums1 {
			minNums2 = nums2[i]
		}
	}

	return minNums2 - minNums2
}
