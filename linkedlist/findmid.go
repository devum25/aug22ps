package linkedlist

func FindMidNodeSecondMid(head *Node) *Node { // in case of even node this will return (n/2)+1 node
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func FindMidNodeFirstMid(head *Node) *Node {
	slow := head
	fast := head
	prev := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		return prev
	}
	return slow
}
