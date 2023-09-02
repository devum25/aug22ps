package linkedlist

func DeleteAtStart(head *Node) *Node {
	head = head.Next
	return head
}

func DeleteAtEnd(head *Node) *Node {
	headNode := head
	for headNode.Next.Next != nil {
		headNode = headNode.Next
	}

	headNode.Next = nil

	return head
}

func DeleteAtKthPosition(head *Node, k int) *Node { //2
	current := head

	if k == 0 {
		head = current.Next
		return head
	}

	for i := 0; i < k-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	return head
}
