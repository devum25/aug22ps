package binarytree

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs1(root, 0)
}

func dfs1(node *TreeNode, current int) int {
	if node == nil {
		return 0
	}
	// Update the current number formed by adding the current node's value
	current = current*10 + node.Val
	// If it's a leaf node, return the current number
	if node.Left == nil && node.Right == nil {
		return current
	}
	// Recursively traverse left and right subtrees, passing the updated current number
	return dfs1(node.Left, current) + dfs1(node.Right, current)
}
