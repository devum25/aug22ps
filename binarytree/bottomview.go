package binarytree

import "sort"

func BottomView(root *TreeNode) []int {
	hash := make(map[int]int)

	hash = helperBottom(root, hash)

	dists := make([]int, 0)
	for dist := range hash {
		dists = append(dists, dist)
	}
	sort.Ints(dists)

	res := make([]int, len(dists))
	for i, dist := range dists {
		res[i] = hash[dist]
	}

	return res
}

type CustomNode struct {
	Node *TreeNode
	Dist int
}

func helperBottom(root *TreeNode, hash map[int]int) map[int]int {

	queue := make([]*CustomNode, 0)

	queue = append(queue, &CustomNode{Dist: 0, Node: root})

	hash[0] = root.Val

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Node.Left != nil {
				x := temp.Dist - 1
				queue = append(queue, &CustomNode{Node: temp.Node.Left, Dist: x})
				hash[x] = temp.Node.Left.Val
			}

			if temp.Node.Right != nil {
				x := temp.Dist + 1
				queue = append(queue, &CustomNode{Node: temp.Node.Right, Dist: x})
				hash[x] = temp.Node.Right.Val
			}
		}
	}

	return hash
}
