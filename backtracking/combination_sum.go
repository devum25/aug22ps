package backtracking

// Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
// frequency of at least one of the chosen numbers is different.
// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

// Example 1:

// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

func CombinationSum(candidates []int, target int) [][]int {
	lst := make([]int, 0)
	ans := make([][]int, 0)
	cuSum := 0

	combinationSum(candidates, target, 0, &cuSum, &lst, &ans)
	return ans
}

func combinationSum(arr []int, k int, idx int, currSum *int, lst *[]int, ans *[][]int) {
	if *currSum > k {
		return
	}
	if *currSum == k {
		// build answer
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
		*currSum = *currSum + arr[i]
		*lst = append(*lst, arr[i])
		combinationSum(arr, k, i, currSum, lst, ans)
		*currSum = *currSum - arr[i]
		*lst = (*lst)[:len(*lst)-1]
	}
}
