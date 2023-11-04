package binarytree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	lst := diameterOfBinaryTree(root.Left)
	lrt := diameterOfBinaryTree(root.Right)

	hl := maxDepth(root.Left)
	hr := maxDepth(root.Right)

	return max(lst, lrt, (hl + hr + 2))
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
type NodeInfo struct {
	height   int
	diameter int
}

func diameterOfBinaryTreeOptimal(root *TreeNode) int {
	i := diameter(root)
	return i.diameter
}

func diameter(root *TreeNode) *NodeInfo {
	if root == nil {
		return &NodeInfo{height: -1, diameter: -1}
	}

	left := diameter(root.Left)
	right := diameter(root.Right)

	return &NodeInfo{height: max1(left.height, right.height) + 1, diameter: max(left.diameter, right.diameter, left.height+right.height+2)}
}

func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
