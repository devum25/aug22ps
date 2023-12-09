package backtracking

import (
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.

// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// Example 2:

// Input: candidates = [2,5,2,1,2], target = 5
// Output:
// [
// [1,2,2],
// [5]
// ]

func CombinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	lst := make([]int, 0)
	sum := 0
	sort.Ints(candidates)
	sum2(candidates, target, 0, &lst, &ans, &sum)
	return ans
}

func sum2(arr []int, k int, idx int, lst *[]int, ans *[][]int, currSum *int) {
	if *currSum > k {
		return
	}
	if *currSum == k {
		// build ans
		x := make([]int, len(*lst))
		for i := 0; i < len(*lst); i++ {
			x[i] = (*lst)[i]
		}
		*ans = append(*ans, x)
		return
	}
	if idx == len(arr) {
		return
	}

	for i := idx; i < len(arr); i++ {
		if arr[i] > k {
			continue
		}
		if i > idx && arr[i] == arr[i-1] {
			continue
		}
		*currSum = *currSum + arr[i]
		*lst = append(*lst, arr[i])
		sum2(arr, k, i+1, lst, ans, currSum)
		*currSum = *currSum - arr[i]
		*lst = (*lst)[:len(*lst)-1]
	}

}
