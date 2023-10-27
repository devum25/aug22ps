package binarytree

func Size(root *TreeNode) int {
	return help(root, 0)
}

func help(root *TreeNode, count int) int {
	if root == nil {
		return count
	}

	count++
	count = help(root.Left, count)
	count = help(root.Right, count)

	return count
}
