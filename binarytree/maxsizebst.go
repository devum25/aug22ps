package binarytree

import "math"

func MaxSize(root *TreeNode) int {
	max := math.MinInt
	helpMaxSize(root, &max)
	return max
}

type SizeInfo struct {
	IsBST bool
	Size  int
	Min   int
	Max   int
}

func helpMaxSize(root *TreeNode, max *int) *SizeInfo {
	if root == nil {
		return &SizeInfo{IsBST: true, Size: 0, Min: math.MaxInt, Max: math.MinInt}
	}

	l := helpMaxSize(root.Left, max)
	r := helpMaxSize(root.Right, max)
	sum := l.Size + r.Size + 1
	if l.IsBST && r.IsBST && l.Max < root.Val && r.Min > root.Val {

		if sum > *max {
			*max = sum
		}
		min := 0
		max := 0
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
		return &SizeInfo{IsBST: true, Size: sum, Min: min, Max: max}
	} else {
		return &SizeInfo{IsBST: false, Size: sum, Min: 0, Max: 0}
	}

}
