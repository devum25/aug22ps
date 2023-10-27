package binarytree

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	res = preorder(root, res)
	return res
}

func preorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	// math.MinInt
	// x := int(math.Max(float64(1), float64(2)))
	res = append(res, root.Val)
	res = preorder(root.Left, res)
	res = preorder(root.Right, res)
	return res
}
