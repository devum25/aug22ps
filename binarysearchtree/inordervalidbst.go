package binarysearchtree

import "github.com/devum25/techbench/binarytree"

func IsValidBSTUsinInOrder(root *binarytree.TreeNode) bool {
	x := helpInOrder(root)
	return check(x)
}

func helpInOrder(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	val1 := helpInOrder(root.Left)
	if val1 != nil {
		res = append(res, val1...)
	}

	res = append(res, root.Val)
	val2 := helpInOrder(root.Right)
	if val2 != nil {
		res = append(res, val2...)
	}

	return res
}

func check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}
