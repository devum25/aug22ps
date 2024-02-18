package linkedlist

func FirstNodeOfCycle(node *MyLinkedList) *MyLinkedList {
	slow := node
	fast := node

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return getfirstNode(node, slow)
		}
	}

	return nil
}

func getfirstNode(head *MyLinkedList, prt *MyLinkedList) *MyLinkedList {
	curr := head
	p := prt

	for curr != p {
		curr = curr.Next
		p = p.Next
	}

	return curr
}
