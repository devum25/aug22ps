package backtracking

// 78. Subsets
// Given an integer array nums of unique elements, return all possible
// subsets
//  (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.
// Example 1:

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

func GenerateSubSet(arr []int) [][]int {
	lst := make([]int, 0)
	ans := make([][]int, 0)
	generateSubset(arr, 0, &lst, &ans)
	return ans
}

func generateSubset(arr []int, currIdx int, lst *[]int, ans *[][]int) {
	if currIdx == len(arr) {
		x := make([]int, len(*lst))
		for i := 0; i < len(x); i++ {
			y := *lst
			x[i] = y[i]
		}
		*ans = append(*ans, x)
		return
	}

	*lst = append(*lst, arr[currIdx])
	generateSubset(arr, currIdx+1, lst, ans)

	y := *lst
	y = y[:len(y)-1]
	*lst = y
	generateSubset(arr, currIdx+1, lst, ans)
}
