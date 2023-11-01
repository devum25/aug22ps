package binarysearchtree

import "github.com/devum25/techbench/binarytree"

func SumOfCommonNode(root1 *binarytree.TreeNode, root2 *binarytree.TreeNode) int {
	s1 := helpInOrder(root1)
	s2 := helpInOrder(root2)
	sum := 0
	if len(s1) > len(s2) {
		for i := 0; i < len(s1); i++ {
			for j := i; j < len(s2); j++ {
				if s1[i] == s2[j] {
					sum += s1[i]
				}
			}
		}
	}

	if len(s2) > len(s1) {
		for i := 0; i < len(s2); i++ {
			for j := i; j < len(s1); j++ {
				if s1[i] == s2[j] {
					sum += s1[i]
				}
			}
		}
	}

	return sum

}
