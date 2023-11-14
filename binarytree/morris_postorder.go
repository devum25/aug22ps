package binarytree

func MorrisPostOrder(root *TreeNode) []int {
	curr := root
	res := make([]int, 0)
	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
		} else {
			tmp := curr.Left

			for tmp.Right != nil && tmp.Right != curr {
				tmp = tmp.Right
			}

			if tmp.Right == nil {
				tmp.Right = curr
				res = append(res, curr.Val)
				curr = curr.Left
			} else if tmp.Right == curr {
				tmp.Right = nil
				curr = curr.Right
			}
		}
	}

	return res
}
