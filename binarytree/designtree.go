package binarytree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}
