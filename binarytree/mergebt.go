package binarytree

func MergeTwoBt(root1 *TreeNode, root2 *TreeNode) {

}

func mergeTrees(r1 *TreeNode, r2 *TreeNode) *TreeNode {
	if r1 == nil && r2 == nil {
		return nil
	}
	val := 0
	if r1 != nil && r2 != nil {
		val = r1.Val + r2.Val
	} else if r1 != nil && r2 == nil {
		val = r1.Val
	} else {
		val = r2.Val
	}

	node := &TreeNode{Val: val}
	if r1 != nil && r2 != nil {
		node.Left = mergeTrees(r1.Left, r2.Left)
		node.Right = mergeTrees(r1.Right, r2.Right)
	} else if r1 != nil && r2 == nil {
		node.Left = mergeTrees(r1.Left, r2)
		node.Right = mergeTrees(r1.Right, r2)
	} else {
		node.Left = mergeTrees(r1, r2.Left)
		node.Right = mergeTrees(r1, r2.Right)
	}

	return node
}

///////////////////////////////////////////////////////////////////////////
func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// if one of t1 and t2 is nil, return the other
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// merge t1 and t2
	root := &TreeNode{Val: t1.Val + t2.Val}
	// recursion
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}
