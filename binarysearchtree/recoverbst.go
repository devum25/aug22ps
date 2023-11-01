package binarysearchtree

func RecoverBST(root *TreeNode) {
	// recover(root)
	info := &Info{}
	recoverOptimal(root, info)
	if info.n1 != nil && info.n2 != nil {
		info.n1.Val, info.n2.Val = info.n2.Val, info.n1.Val
	}
}

type Info struct {
	prev *TreeNode
	n1   *TreeNode
	n2   *TreeNode
}

func recoverOptimal(root *TreeNode, info *Info) {
	if root == nil {
		return
	}

	recoverOptimal(root.Left, &Info{})
	if info.prev != nil {
		if info.prev.Val > root.Val {
			if info.n1 == nil {
				info.n1 = info.prev
				info.n2 = root
			} else {
				info.n2 = root
			}
		}
	}
	info.prev = root
	recoverOptimal(root.Right, info)

}

func recover(root *TreeNode) {
	if root == nil {
		return
	}

	list := inorder(root)
	var prev *TreeNode
	var n1 *TreeNode
	var n2 *TreeNode
	for i := 0; i < len(list); i++ {
		if prev != nil {
			if prev.Val > list[i].Val {
				if n1 == nil {
					n1 = prev
					n2 = list[i]
				} else {
					n2 = list[i]
				}
			}
		}
		prev = list[i]
	}

	n1.Val, n2.Val = n2.Val, n1.Val

}

func inorder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	res := make([]*TreeNode, 0)

	res = append(res, inorder(root.Left)...)
	res = append(res, root)
	res = append(res, inorder(root.Right)...)

	return res
}
