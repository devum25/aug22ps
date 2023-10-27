package binarytree

func BoundaryView(root *TreeNode) []int {

	x := helperBoundary(root)

	ans := make([]int, 0)

	for i := 0; i < len(x); i++ {
		if i == len(x)-1 {
			ans = append(ans, x[i]...)
		} else {
			ans = append(ans, x[i][0])
		}
	}

	for i := len(x) - 2; i >= 1; i-- {
		ans = append(ans, x[i][len(x[i])-1])
	}
	return ans
}

func helperBoundary(root *TreeNode) [][]int {
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
