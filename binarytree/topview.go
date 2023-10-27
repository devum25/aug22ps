package binarytree

import "sort"

func TopView(root *TreeNode) []int {
	hash := make(map[int]int)
	hash = helpTopView(root, hash)

	temp := make([]int, len(hash))
	i := 0
	for k := range hash {
		temp[i] = k
		i++
	}

	sort.Ints(temp)
	newHash := make(map[int]int)
	for _, v := range temp {
		newHash[v] = hash[v]
	}
	res := make([]int, len(hash))
	i = 0
	for _, v := range newHash {
		res[i] = v
	}

	return res
}

type topNode struct {
	Node *TreeNode
	Dist int
}

func helpTopView(root *TreeNode, hash map[int]int) map[int]int {
	if root == nil {
		return hash
	}

	queue := make([]*topNode, 0)

	queue = append(queue, &topNode{Node: root, Dist: 0})
	hash[0] = root.Val

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Node.Left != nil {
				x := node.Dist - 1
				queue = append(queue, &topNode{Node: node.Node.Left, Dist: x})
				_, ok := hash[x]
				if !ok {
					hash[x] = node.Node.Left.Val
				}
			}
			if node.Node.Right != nil {
				x := node.Dist + 1
				queue = append(queue, &topNode{Node: node.Node.Right, Dist: x})
				_, ok := hash[x]
				if !ok {
					hash[x] = node.Node.Right.Val
				}
			}
		}
	}

	return hash

}
