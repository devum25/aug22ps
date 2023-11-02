package binarysearchtree

// Given a binary search tree represented by root A, write a function to find the Bth smallest element in the tree.
// Problem Constraints
// 1 <= Number of nodes in binary tree <= 100000
// 0 <= node values <= 10^9
// Input Format
// First and only argument is head of the binary tree A.
// Output Format
// Return an integer, representing the Bth element.

func KthSmallestElemet(root *TreeNode, k int) int {

	return kthSmallest(root, k)
}

func kthSmallest(root *TreeNode, k int) int {
	asc := inorderK(root)

	return asc[k-1]
}

func inorderK(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	res = append(res, inorderK(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderK(root.Right)...)

	return res
}
