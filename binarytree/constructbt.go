package binarytree

// Q6. Binary Tree From Inorder And Postorder
// Problem Description
// Given the inorder and postorder traversal of a tree, construct the binary tree.
// A = [2, 1, 3] : Inorder
// B = [2, 3, 1] : postorder

//  1
// / \
// 2   3

// A = [6, 1, 3, 2] :inorder
// B = [6, 3, 2, 1] : postorder

//   1
// 	/ \
//     2
//    / \
//   3   6

func BuildTree(A []int, B []int) *TreeNode {
	return getTree(A, B)
}

func getTree(io []int, po []int) *TreeNode {
	var root *TreeNode
	root = tree(root, io, po, 0, len(io)-1, len(io)-1)
	return root
}

func tree(root *TreeNode, io []int, po []int, s, e, pos int) *TreeNode {
	if s > e {
		return nil
	}

	root = &TreeNode{Val: po[pos]}
	idx := getIndex(io, po[pos])
	root.Left = tree(root, io, po, s, idx-1, pos-(e-idx)-1)
	root.Right = tree(root, io, po, idx+1, e, pos-1)
	return root
}

func getIndex(io []int, k int) int {
	for i := 0; i < len(io); i++ {
		if io[i] == k {
			return i
		}
	}
	return -1
}
