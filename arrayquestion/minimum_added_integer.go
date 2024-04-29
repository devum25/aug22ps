package arrayquestion

import (
	"math"
	"sort"
)

func MinimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)

	n := len(nums2)

	x := len(nums1) - n
	temp := math.MaxInt
	for i := 0; i <= x; i++ {
		ans := solve(nums1[i:n+i], nums2)
		if math.Abs(float64(ans)) < math.Abs(float64(temp)) {
			temp = ans
		}
	}

	return temp

}

func solve(nums1 []int, nums2 []int) int {
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums2[i] - nums1[i]
	}
	x := sum / len(nums1)
	return x
}
