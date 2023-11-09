package binarytree

// 116. Populating Next Right Pointers in Each Node

// You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

// Initially, all next pointers are set to NULL.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	return nil
}

func bfs(root *Node) {
	q := make([]*Node, 0)

	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		var prev *Node
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if prev != nil {
				prev.Next = node
				prev = node
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			if prev == nil {
				prev = node
			}
		}
	}

}
