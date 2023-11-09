package binarytree

func TargetSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return TargetSum(root.Left, sum-root.Val) || TargetSum(root.Right, sum-root.Val)
}
