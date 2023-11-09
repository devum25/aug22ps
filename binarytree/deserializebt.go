package binarytree

func DeserializeBt(arr []int) *TreeNode {
	return deserialize_levelorder(arr)
}

func deserialize_levelorder(arr []int) *TreeNode {
	root := &TreeNode{Val: arr[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	end := len(arr) - 1
	i := 1
	for i <= end && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		j := 0
		for j < 2 {
			if j == 0 {
				if arr[i] == -1 {
					node.Left = nil
				} else {
					node.Left = &TreeNode{Val: arr[i]}
					queue = append(queue, node.Left)
				}
				i++
				j++
			} else {
				if arr[i] == -1 {
					node.Right = nil
				} else {
					node.Right = &TreeNode{Val: arr[i]}
					queue = append(queue, node.Right)
				}
				i++
				j++
			}
		}

	}

	return root

}
