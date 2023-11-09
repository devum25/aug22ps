package binarytree

func NextPointer(root *Node) {
	nextpointer(root)
}

func nextpointer(root *Node) {
	if root == nil {
		return
	}

	level := root

	for level != nil {
		curr := level
		for curr != nil {
			if curr.Left != nil {
				if curr.Right != nil {
					curr.Left.Next = curr.Right
				} else {
					right := getNextRightPointer(curr)
					curr.Left.Next = right
				}
			}
			if curr.Right != nil {
				curr.Right.Next = getNextRightPointer(curr)
			}

			curr = curr.Next
		}
		if level.Left != nil {
			level = level.Left
		} else if level.Right != nil {
			level = level.Right
		} else {
			level = getNextRightPointer(level)
		}

	}

}

func getNextRightPointer(root *Node) *Node {
	if root == nil {
		return nil
	}

	temp := root.Next

	for temp != nil {
		if temp.Left != nil {
			return temp.Left
		}
		if temp.Right != nil {
			return temp.Right
		}

		temp = temp.Next
	}

	return nil
}
