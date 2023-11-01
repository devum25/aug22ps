package binarysearchtree

func TrimBinarySearchTree(root *TreeNode, low, high int) *TreeNode {
	return trim(root, low, high)
}

func trim(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trim(root.Right, low, high)
	} else if root.Val > high {
		return trim(root.Left, low, high)
	} else {
		root.Left = trim(root.Left, low, high)
		root.Right = trim(root.Right, low, high)
	}

	return root
}
