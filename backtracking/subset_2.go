package backtracking

import "sort"

// Given a set of distinct integers A, return all possible subsets.

// NOTE:

// Elements in a subset must be in non-descending order. // ascending
// The solution set must not contain duplicate subsets. // unique
// Also, the subsets should be sorted in ascending ( lexicographic ) order.
// The initial list is not necessarily sorted.

// [1, 2, 3]
// [
//  []
//  [1]
//  [1, 2]
//  [1, 2, 3]
//  [1, 3]
//  [2]
//  [2, 3]
//  [3]
// ]

// func Subsets(A []int) [][]int {
// 	ans := make([][]int, 0)
// 	// x := make([]int, 0)
// 	// ans = append(ans, x)
// 	lst := make([]int, 0)
// 	sort.Ints(A)
// 	subsetofarray(A, 0, &lst, &ans)
// 	sort.Slice(ans, func(i, j int) bool {
// 		return ans[i][j] < ans[i][j]
// 	})

// 	return ans
// }

// func subsetofarray(arr []int, idx int, lst *[]int, ans *[][]int) {
// 	if idx == len(arr) {
// 		x := make([]int, len(*lst))
// 		temp := *lst
// 		for i := 0; i < len(*lst); i++ {
// 			x[i] = temp[i]
// 		}
// 		*ans = append(*ans, x)
// 		return
// 	}

// 	*lst = append(*lst, arr[idx])
// 	subsetofarray(arr, idx+1, lst, ans)

// 	*lst = (*lst)[:len(*lst)-1]
// 	subsetofarray(arr, idx+1, lst, ans)
// }

// var a []int
// var ans [][]int

// func subset(temp []int, index int) {
// 	// insert the current subset to our answer
// 	ans = append(ans, temp)
// 	for i := index; i < len(a); i++ {
// 		// for every element we can either take it or skip it
// 		temp = append(temp, a[i])
// 		subset(temp, i+1)
// 		var cur []int
// 		for i := 0; i < len(temp)-1; i++ {
// 			cur = append(cur, temp[i])
// 		}
// 		temp = cur
// 	}
// 	return
// }
// func subsets(A []int) [][]int {
// 	ans = nil
// 	sort.Ints(A)
// 	a = A
// 	index := 0
// 	var temp []int
// 	subset(temp, index)
// 	return ans
// }

func GenerateSubSet_Duplicate(arr []int) [][]int {
	lst := make([]int, 0)
	ans := make([][]int, 0)
	sort.Ints(arr)
	generateSubset(arr, 0, &lst, &ans)
	return ans
}

func generateSubset_duplicate(arr []int, currIdx int, lst *[]int, ans *[][]int) {
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
	generateSubset_duplicate(arr, currIdx+1, lst, ans)

	y := *lst
	y = y[:len(y)-1]
	*lst = y

	for (currIdx+1) < len(arr) && arr[currIdx] == arr[currIdx+1] {
		currIdx++
	}
	generateSubset_duplicate(arr, currIdx+1, lst, ans)
}
