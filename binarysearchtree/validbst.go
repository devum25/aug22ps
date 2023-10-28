package binarysearchtree

import (
	"math"

	"github.com/devum25/techbench/binarytree"
)

func IsValidBST(root *binarytree.TreeNode) bool {
	return false
}

type NodeInfo struct {
	IsValid bool
	Max     int
	Min     int
}

func isBST(root *binarytree.TreeNode) *NodeInfo {
	if root == nil {
		return &NodeInfo{IsValid: true, Max: math.MinInt, Min: math.MaxInt}
	}

	n1 := isBST(root.Left)
	n2 := isBST(root.Right)

	if (n1.IsValid && n2.IsValid) && (n1.Max < root.Val) && (n2.Min > root.Val) {
		return &NodeInfo{IsValid: true, Max: int(math.Max(float64(root.Val), float64(n2.Max))),
			Min: int(math.Min(float64(root.Val), float64(n2.Min)))}
	}

	return &NodeInfo{IsValid: false, Max: int(math.Max(float64(root.Val), float64(n2.Max))),
		Min: int(math.Min(float64(root.Val), float64(n2.Min)))}
}
