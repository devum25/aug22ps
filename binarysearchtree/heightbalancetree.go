package binarysearchtree

import (
	"math"

	"github.com/devum25/techbench/binarytree"
)

func IsHeightBalanced(root *binarytree.TreeNode) bool {
	return helpHBT(root)
}

func helpHBT(root *binarytree.TreeNode) bool {
	if root == nil {
		return true
	}

	h1 := height(root.Left)
	h2 := height(root.Right)

	if int(math.Abs(float64(h1)-float64(h2))) > 1 {
		return false
	}

	return helpHBT(root.Left) && helpHBT(root.Right)

}

func height(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}

	h1 := height(root.Left)
	h2 := height(root.Right)

	return int(math.Max(float64(h1), float64(h2))) + 1
}
