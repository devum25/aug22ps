package binarysearchtree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helpLCA(root, p, q)
}

func helpLCA(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		return helpLCA(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return helpLCA(root.Left, p, q)
	} else {
		return root
	}

}

//  optimal way to do it
func LCA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := LCA(root.Left, p, q)
	r := LCA(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	} else {
		return r
	}

}
