package binarysearchtree

import "github.com/devum25/techbench/binarytree"

func InsertInBST(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil {
		return &binarytree.TreeNode{Val: val}
	}
	return helpInsert(root, val)
}

func helpInsert(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil {
		return nil
	}
	var left *binarytree.TreeNode
	var right *binarytree.TreeNode
	if val < root.Val {
		left = helpInsert(root.Left, val)
		if left == nil {
			root.Left = &binarytree.TreeNode{Val: val}
		}
	}

	if val > root.Val {
		right = helpInsert(root.Right, val)
		if right == nil {
			root.Right = &binarytree.TreeNode{Val: val}
		}
	}

	return root
}
