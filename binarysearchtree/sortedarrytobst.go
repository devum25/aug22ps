package binarysearchtree

import "github.com/devum25/techbench/binarytree"

func SortedArrayToBST(nums []int) *binarytree.TreeNode {
	mid := (len(nums) - 1) / 2

	root := &binarytree.TreeNode{Val: nums[mid]}
	root.Left = helpBST(nums, 0, mid-1)
	root.Right = helpBST(nums, mid+1, len(nums)-1)
	return root
}

func helpBST(nums []int, s, e int) *binarytree.TreeNode {
	if s > e {
		return nil
	}

	mid := (s + e) / 2

	root := &binarytree.TreeNode{Val: nums[mid]}
	root.Left = helpBST(nums, s, mid-1)
	root.Right = helpBST(nums, mid+1, e)

	return root
}
