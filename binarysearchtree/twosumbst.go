package binarysearchtree

func TwoSum(root *TreeNode, k int) int {
	hash := make(map[int]int)

	x := find(root, k, hash)
	if x {
		return 1
	}
	return 0
}

func find(root *TreeNode, k int, set map[int]int) bool {
	if root == nil {
		return false
	}

	if _, ok := set[k-root.Val]; ok {
		return true
	}

	set[root.Val] = root.Val
	if root.Left != nil {
		if _, ok := set[k-root.Left.Val]; ok {
			return true
		}
	}

	if root.Right != nil {
		if _, ok := set[k-root.Right.Val]; ok {
			return true
		}
	}
	x := find(root.Left, k, set)
	y := find(root.Right, k, set)

	return x || y
}

////////////////////////////////////////////////////////////

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func dfs(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}

	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true

	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}
