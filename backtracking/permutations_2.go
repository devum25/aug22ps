package backtracking

import (
	"sort"
)

// 47. Permutations II
// Medium
// 8.2K
// 138
// Companies
// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

// Example 1:
// Input: nums = [1,1,2]
// Output:
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]
// Example 2:

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func PermuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	permute(nums, 0, &ans)
	return ans
}

func permute(arr []int, idx int, ans *[][]int) {
	if len(arr) == idx {
		temp := make([]int, len(arr))

		for i := 0; i < len(temp); i++ {
			temp[i] = arr[i]
		}
		*ans = append(*ans, temp)
		return
	}
	hash := make(map[int]bool)
	for i := idx; i < len(arr); i++ {
		if _, ok := hash[arr[i]]; ok {
			continue
		}
		hash[arr[i]] = true
		arr[i], arr[idx] = arr[idx], arr[i]
		permute(arr, idx+1, ans)
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
