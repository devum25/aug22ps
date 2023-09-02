package linkedlist

type Node struct {
	Val  int
	Next *Node
}

func InsertAtStart(head *Node, val int) *Node {
	newNode := &Node{Val: val, Next: head}
	return newNode
}

func InsertAtEnd(head *Node, val int) {
	headNode := head
	for headNode.Next != nil {
		headNode = headNode.Next
	}

	headNode.Next = &Node{Val: val}
}

func InsertAtKthPosition(head *Node, k, val int) *Node { //2
	current := head
	newNode := &Node{Val: val}

	if k == 0 {
		newNode.Next = current
		return newNode
	}

	for i := 0; i < k-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	return head
}
