package binarytree

func ZigzagLevelOrder(root *TreeNode) [][]int {
	x := helperzigzag(root)

	ans := make([][]int, len(x))

	for i := 0; i < len(x); i++ {
		if i == 0 {
			ans[0] = append(ans[0], x[i][0])
		} else if i == 1 || i%2 != 0 {
			for j := len(x[i]) - 1; j >= 0; j-- {
				ans[i] = append(ans[i], x[i][j])
			}
		} else if i%2 == 0 {
			for j := 0; j < len(x[i]); j++ {
				ans[i] = append(ans[i], x[i][j])
			}
		}
	}

	return ans
}

func helperzigzag(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		list := make([]int, 0)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			list = append(list, temp.Val)

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}

		}

		res = append(res, list)
	}

	return res
}
