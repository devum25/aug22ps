package backtracking

import (
	"sort"
)

// You are given an array nums of positive integers and a positive integer k.
// A subset of nums is beautiful if it does not contain two integers with an absolute difference equal to k.
// Return the number of non-empty beautiful subsets of the array nums.
// A subset of nums is an array that can be obtained by deleting some (possibly none) elements from nums. Two subsets are different if and only if the chosen indices to delete are different.

func BeautifulSubsets(nums []int, k int) int {
	ans := [][]int{}
	lst := []int{}
	hash := make(map[int]int, 1000)

	for i := 1; i <= 1000; i++ {
		hash[i] = 0
	}

	sort.Ints(nums)
	solve2(nums, lst, &ans, 0, k, hash)

	return len(ans) - 1
}

func solve2(nums []int, lst []int, ans *[][]int, idx int, k int, hash map[int]int) {
	if idx == len(nums) {
		temp := make([]int, len(lst))

		for i := 0; i < len(temp); i++ {
			temp[i] = lst[i]
		}
		*ans = append(*ans, temp)
		return
	}
	if hash[nums[idx]] == 0 {
		lst = append(lst, nums[idx])
		hash[nums[idx]-k]++
		hash[nums[idx]+k]++
		solve2(nums, lst, ans, idx+1, k, hash)
		hash[lst[len(lst)-1]-k]--
		hash[lst[(len(lst)-1)]+k]--
		lst = lst[:len(lst)-1]
	}

	solve2(nums, lst, ans, idx+1, k, hash)
}
