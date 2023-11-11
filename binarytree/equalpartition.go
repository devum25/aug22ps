package binarytree

func EqualPartition(root *TreeNode) bool {
	hash := make(map[int]bool)
	val := partition(root, hash)
	if _, ok := hash[val/2]; ok {
		return true
	}
	return false
}

func partition(root *TreeNode, hash map[int]bool) int {
	if root == nil {
		return 0
	}

	val1 := partition(root.Left, hash)
	val2 := partition(root.Right, hash)

	val := val1 + val2 + root.Val

	hash[val] = true

	return val

}
