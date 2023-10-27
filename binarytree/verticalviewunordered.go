package binarytree

func VerticalViewUnordered(root *TreeNode) [][]int {
	hash := make(map[int][]*TreeNode)
	hash = helpView(root, 0, hash)
	res := make([][]int, 0)
	for _, v := range hash {
		temp := make([]int, 0)
		for i := 0; i < len(v); i++ {
			temp = append(temp, v[i].Val)
		}
		res = append(res, temp)
	}

	return res
}

func helpView(root *TreeNode, dist int, hash map[int][]*TreeNode) map[int][]*TreeNode {
	if root == nil {
		return nil
	}

	if hash[dist] == nil {
		hash[dist] = []*TreeNode{root}
	} else {
		hash[dist] = append(hash[dist], root)
	}

	helpView(root.Left, dist-1, hash)
	helpView(root.Right, dist+1, hash)

	return hash
}
