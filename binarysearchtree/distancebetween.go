package binarysearchtree

func solve(root *TreeNode, B int, C int) int {
	lca := LCA(root, B, C)
	left := dist(lca, B)
	right := dist(lca, C)

	return left + right
}

func dist(root *TreeNode, B int) int {
	if root.Val == B {
		return 0
	}

	if B < root.Val {
		return 1 + dist(root.Left, B)
	}
	return 1 + dist(root.Right, B)
}

func LCA(root *TreeNode, B, C int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val > B && root.Val > C {
		return LCA(root.Left, B, C)
	} else if root.Val < B && root.Val < C {
		return LCA(root.Right, B, C)
	} else {
		return root
	}

}
