package binarytree

import "fmt"

func MorrisIteration(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left == nil {
			fmt.Print(curr.Val)
			curr = curr.Right
		} else {
			temp := curr.Left

			for temp.Right != nil && temp.Right != curr {
				temp = temp.Right
			}

			if temp.Right == nil {
				temp.Right = curr
				curr = curr.Left
			} else if temp.Right == curr {
				temp.Right = nil
				fmt.Print(curr.Val)
				curr = curr.Right
			}
		}
	}
}
