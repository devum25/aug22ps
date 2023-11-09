package binarytree

import "math"

func MaxSumBT(root *TreeNode) int {
	max := math.MinInt
	dfs(root, &max)
	return max
}

type Valid struct {
	isBST bool
	sum   int
	min   int
	max   int
}

func dfs(root *TreeNode, max *int) *Valid {
	if root == nil {
		return &Valid{isBST: true, sum: 0, min: math.MaxInt, max: math.MinInt}
	}

	l := dfs(root.Left, max)
	r := dfs(root.Right, max)

	if l.isBST && r.isBST && l.max < root.Val && root.Val < r.min {
		sum := l.sum + r.sum + root.Val
		if sum > *max {
			*max = sum
		}
		var min, max int
		if root.Left != nil {
			min = root.Left.Val
		} else {
			min = root.Val
		}

		if root.Right != nil {
			max = root.Right.Val
		} else {
			max = root.Val
		}
		return &Valid{isBST: true, sum: sum, min: min, max: max}
	} else {
		return &Valid{isBST: false, sum: 0, min: 0, max: 0}
	}

}
