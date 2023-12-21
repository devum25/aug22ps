package backtracking

// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

// You may return the answer in any order.

// Example 1:

// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
// Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
// Example 2:

// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is 1 choose 1 = 1 total combination.
func Combination(n int, k int) [][]int {
	ans := make([][]int, 0)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	lst := make([]int, 0)

	generatecomb(arr, 0, &lst, &ans, k)

	return ans
}

func generatecomb(arr []int, idx int, lst *[]int, ans *[][]int, k int) {
	if len(*lst) == k {
		temp := make([]int, k)

		for i := 0; i < k; i++ {
			temp[i] = (*lst)[i]
		}
		*ans = append(*ans, temp)
		return
		// add list to answer
	}

	if idx == len(arr) {
		return
	}

	*lst = append(*lst, arr[idx])
	generatecomb(arr, idx+1, lst, ans, k)

	*lst = (*lst)[:len(*lst)-1]
	generatecomb(arr, idx+1, lst, ans, k)
}
