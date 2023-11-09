package binarytree

func SerializeBT(root *TreeNode) []int {
	return serailize(root)
}

func serailize(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 0)

	q = append(q, root)

	ans := make([]int, 0)
	ans = append(ans, root.Val)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				ans = append(ans, node.Left.Val)
				q = append(q, node.Left)
			} else {
				ans = append(ans, -1)
			}

			if node.Right != nil {
				ans = append(ans, node.Right.Val)
				q = append(q, node.Right)
			} else {
				ans = append(ans, -1)
			}
		}

	}

	return ans
}
