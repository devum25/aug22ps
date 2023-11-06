package binarytree

func BoundaryView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	leaves := []int{}
	left := getLeft(root.Left)
	right := getRight(root.Right)
	res = append(res, left[:len(left)-1]...)
	getLeaves(root, &leaves)
	res = append(res, leaves...)
	res = append(res, reverse(right[:len(right)-1])...)

	return res
}

func getLeft(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	x := make([]int, 0)

	if root.Left != nil {

		// to ensure top down order, print the node
		// before calling itself for left subtree
		x = append(x, root.Val)
		getLeft(root.Left)
	} else if root.Right != nil {
		x = append(x, root.Val)
		getLeft(root.Right)
	}

	return x
}

func getRight(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	x := make([]int, 0)

	if root.Right != nil {
		x = append(x, root.Val)
		getRight(root.Right)
	} else if root.Left != nil {
		x = append(x, root.Val)
		getRight(root.Left)
	}

	return x
}

func getLeaves(root *TreeNode, x *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*x = append(*x, root.Val)
	}

	getLeaves(root.Left, x)
	getLeaves(root.Right, x)
}

func reverse(x []int) []int {
	i := 0
	j := len(x) - 1

	for i < j {
		x[i], x[j] = x[j], x[i]
	}

	return x
}
