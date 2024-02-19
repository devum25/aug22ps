package arrayquestion

import "sort"

// Given an array, A of integers of size N. Find the maximum value of j - i such that A[i] <= A[j].
// Return an integer denoting the maximum value of j - i.
func MaximumGap(A []int) int {
	lst := make([]Node, 0)
	for i := 0; i < len(A); i++ {
		lst = append(lst, Node{Val: A[i], Idx: i})
	}

	sort.SliceStable(lst, func(i, j int) bool {
		return lst[i].Val < lst[j].Val
	})
	ans := 0
	maxIdx := lst[len(lst)-1].Idx
	for i := len(lst) - 2; i >= 0; i-- {
		if maxIdx-lst[i].Idx > 0 && maxIdx-lst[i].Idx > ans {
			ans = maxIdx - lst[i].Idx
		}

		if lst[i].Idx > maxIdx {
			maxIdx = lst[i].Idx
		}
	}

	return ans
}

type Node struct {
	Val int
	Idx int
}
