package binarysearchtree

import "github.com/devum25/techbench/binarytree"

func RangeSumBST(root *binarytree.TreeNode, low int, high int) int {

	x := helpInOrder1(root)
	sum := 0
	lowest := false
	for i := 0; i < len(x); i++ {
		if x[i] == low {
			sum = sum + x[i]
			lowest = true
		} else if lowest {
			sum = sum + x[i]
		} else if x[i] == high {
			sum = sum + x[i]
			lowest = false
			break
		}
	}

	return sum
}

func helpInOrder1(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	val1 := helpInOrder1(root.Left)
	if val1 != nil {
		res = append(res, val1...)
	}

	res = append(res, root.Val)

	val2 := helpInOrder1(root.Right)
	if val2 != nil {
		res = append(res, val2...)
	}

	return res
}
