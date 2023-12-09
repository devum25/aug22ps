package backtracking

// Given an array nums of distinct integers, return all the possible permutations.
// You can return the answer in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:

// Input: nums = [0,1]
// Output: [[0,1],[1,0]]
// Example 3:

// Input: nums = [1]
// Output: [[1]]

func Permute(nums []int) [][]int {
	ans := make([][]int, 0)
	permutation(nums, &ans, 0)
	return ans
}

func permutation(arr []int, ans *[][]int, idx int) {
	if len(arr) == idx {
		x := make([]int, len(arr))
		for i := 0; i < len(x); i++ {
			x[i] = arr[i]
		}
		*ans = append(*ans, x)
		return
	}

	for i := idx; i < len(arr); i++ {
		arr[idx], arr[i] = arr[i], arr[idx]
		permutation(arr, ans, idx+1)
		arr[i], arr[idx] = arr[idx], arr[i]
	}

}
