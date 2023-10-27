package binarytree

import "sort"

func VerticalViewOrdered(root *TreeNode) [][]int {
	hash := make(map[int][]int)
	helpTraversal(root, 0, hash)

	// Extract and sort the column positions
	columns := make([]int, 0)
	for col := range hash {
		columns = append(columns, col)
	}
	sort.Ints(columns)

	// Create a new map with sorted keys
	newHash := make(map[int][]int)
	for _, col := range columns {
		newHash[col] = hash[col]
	}

	// Construct the final result
	result := make([][]int, len(newHash))
	for i, col := range columns {
		result[i] = make([]int, 0)
		for _, node := range newHash[col] {
			result[i] = append(result[i], node)
		}
	}

	return result
}

type queueNode struct {
	Node *TreeNode
	Dist int
}

func helpTraversal(root *TreeNode, dist int, hash map[int][]int) map[int][]int {
	if root == nil {
		return hash
	}

	queue := make([]*queueNode, 0)

	queue = append(queue, &queueNode{Node: root, Dist: 0})

	if hash[queue[0].Dist] == nil {
		hash[queue[0].Dist] = []int{root.Val}
	}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Node.Left != nil {
				x := temp.Dist - 1
				queue = append(queue, &queueNode{Node: temp.Node.Left, Dist: x})

				if hash[x] == nil {
					hash[x] = []int{temp.Node.Left.Val}
				} else {
					hash[x] = append(hash[x], temp.Node.Left.Val)
				}
			}
			if temp.Node.Right != nil {
				x := temp.Dist + 1
				queue = append(queue, &queueNode{Node: temp.Node.Right, Dist: x})
				if hash[x] == nil {
					hash[x] = []int{temp.Node.Right.Val}
				} else {
					hash[x] = append(hash[x], temp.Node.Right.Val)
				}
			}
		}
	}

	return hash

}
