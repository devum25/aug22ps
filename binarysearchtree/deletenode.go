package binarysearchtree

import (
	"math"
)

func DeleteNode(root *TreeNode, k int) *TreeNode {
	return delete(root, k)
}

func delete(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > k {
		root.Left = delete(root.Left, k)
	} else if root.Val < k {
		root.Right = delete(root.Right, k)
	} else {
		if isLeaf(root) {
			return nil
		} else if ishavingonechild(root) {
			if root.Left != nil {
				return root.Left
			} else {
				return root.Right
			}
		} else {
			max := getMax(root.Left)
			root.Val = max
			root.Left = delete(root.Left, max)
			return root
		}
	}

	return root
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func ishavingonechild(root *TreeNode) bool {
	if root.Right != nil && root.Left == nil {
		return true
	} else if root.Left != nil && root.Right == nil {
		return true
	} else {
		return false
	}
}

func getMax(root *TreeNode) int {
	if root == nil {
		// Handle the case where the root is nil (empty tree).
		return math.MinInt
	}

	// Initialize max to the value of the current node.
	max := root.Val

	// Recursively find the maximum value in the right subtree.
	if root.Right != nil {
		x := getMax(root.Right)
		if x > max {
			max = x
		}
	}

	// Return the maximum value found in the current subtree.
	return max
}
