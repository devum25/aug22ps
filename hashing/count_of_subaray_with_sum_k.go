package hashset

//Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.
import "math"

func SubarraySum(nums []int, k int) int {
	pf := make([]int, len(nums))

	pf[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		pf[i] = pf[i-1] + nums[i]
	}
	count := 0
	for s := 0; s < len(pf); s++ {
		for e := s; e < len(pf); e++ {
			if s == 0 {
				if pf[e] == k {
					count++
				}
			} else {
				if (pf[e] - pf[s-1]) == k {
					count++
				}
			}
		}
	}

	return count
}

func SubarraySumOConstant(nums []int, k int) int {
	pf := make([]int, len(nums))
	hash := make(map[int]int)
	pf[0] = nums[0]
	hash[pf[0]]++
	count := 0
	for i := 1; i < len(nums); i++ {
		pf[i] = pf[i-1] + nums[i]
		hash[pf[i]]++
	}

	for i := 0; i < len(nums); i++ {
		x := pf[i]
		if v, ok := hash[int(math.Abs(float64(x-k)))]; ok {
			count += v
		}
	}

	return count
}
