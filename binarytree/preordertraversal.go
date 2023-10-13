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

	res = append(res, root.val)
	res = preorder(root.left, res)
	res = preorder(root.right, res)
	return res
}
