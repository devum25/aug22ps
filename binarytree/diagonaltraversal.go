package binarytree

import "sort"

// Problem Description

// Consider lines of slope -1 passing between nodes.

// Given a Binary Tree A containing N nodes, return all diagonal elements in a binary tree belonging to same line.

// NOTE:

// See Sample Explanation for better understanding.
// Order does matter in the output.
// To get the same order as in the output traverse the tree same as we do in pre-order traversal.
// Problem Constraints
// 0 <= N <= 105
// Input Format
// First and only Argument represents the root of binary tree A.
// Output Format
// Return a 1D array denoting the diagonal traversal of the tree.

func DiagonalTraversal(root *TreeNode) []int {
	hash := helpDiagonal(root, 0, map[int][]int{})

	res := make([]int, 0)

	cols := make([]int, len(hash))
	i := 0
	for k, _ := range hash {
		cols[i] = k
		i++
	}

	sort.Ints(cols)

	for _, col := range cols {
		res = append(res, hash[col]...)
	}

	return res
}

func helpDiagonal(root *TreeNode, slope int, hash map[int][]int) map[int][]int {
	if root == nil {
		return hash
	}

	if _, ok := hash[slope]; !ok {
		hash[slope] = []int{root.Val}
	} else {
		hash[slope] = append(hash[slope], root.Val)
	}

	hash = helpDiagonal(root.Left, slope+1, hash)
	hash = helpDiagonal(root.Right, slope, hash)

	return hash

}
