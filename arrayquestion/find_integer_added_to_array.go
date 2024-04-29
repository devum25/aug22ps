package arrayquestion

import (
	"math"
	"sort"
)

func MinimumAddedInteger2(nums1 []int, nums2 []int) int {

	sort.Ints(nums1)

	sort.Ints(nums2)

	ans := math.MaxInt

	for i := 0; i < len(nums1); i++ {
		diff := nums2[0] - nums1[i]

		if check(nums1, nums2, diff) {
			ans = min(ans, diff)
		}
	}

	return ans
}

func check(nums1, nums2 []int, diff int) bool {
	count := 0
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i]+diff != nums2[j] {
			count++
		} else {
			j++
		}
		i++
	}
	return (count <= 2) && (j == len(nums2))
}
