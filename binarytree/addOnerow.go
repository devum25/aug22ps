package binarytree

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 && root != nil {
		node := &TreeNode{Val: val}
		node.Left = root
		return node
	}

	if root == nil {
		return nil
	}
	solve4(root, val, depth)
	return root
}

func solve4(root *TreeNode, val, depth int) {
	queue := make([]*TreeNode, 0)
	// prevQueue := make([]*TreeNode, 0)
	queue = append(queue, root)
	c := 0
	done := false
	for len(queue) > 0 && !done {
		c++
		// prevQueue = prevQueue[:0]
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]

			if c == (depth - 1) {
				if temp.Left == nil && temp.Right == nil {
					temp.Left = &TreeNode{Val: val}
					temp.Right = &TreeNode{Val: val}
					done = true
					continue
				}
				if temp.Left != nil {
					left := temp.Left
					temp.Left = &TreeNode{Val: val, Left: left}
				} else {
					temp.Left = &TreeNode{Val: val}
				}

				if temp.Right != nil {
					right := temp.Right
					temp.Right = &TreeNode{Val: val, Right: right}
				} else {
					temp.Right = &TreeNode{Val: val}
				}
				done = true
				continue
			}

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
	}
}
