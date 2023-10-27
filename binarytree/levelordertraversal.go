package binarytree

func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	queue := make([]TreeNode, 0)

	queue = append(queue, *root)

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			level = append(level, temp.Val)
			if temp.Left != nil {
				queue = append(queue, *temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, *temp.Right)
			}
		}

		result = append(result, level)
	}

	return result

}
