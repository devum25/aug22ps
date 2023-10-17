package binarytree

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.left)
	right := maxDepth(root.right)
	return int(math.Max(float64(left), float64(right))) + 1
}
