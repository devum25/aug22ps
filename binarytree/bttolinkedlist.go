package binarytree

func BTToLinkedList(root *TreeNode) {
	var res *TreeNode
	head := &TreeNode{}
	_, head = toPreOrderLinkedList(root, res, head)
	root = head

}

func toPreOrderLinkedList(root *TreeNode, res *TreeNode, head *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return res, head
	}

	if res == nil {
		res = &TreeNode{Val: root.Val}
		head = res
	} else {
		res.Right = &TreeNode{Val: root.Val}
		res = res.Right
	}

	res, head = toPreOrderLinkedList(root.Left, res, head)
	res, head = toPreOrderLinkedList(root.Right, res, head)

	return res, head
}
