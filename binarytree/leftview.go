package binarytree

func LeftView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = append(res, temp.Val)
			}

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}

	}

	return res
}
