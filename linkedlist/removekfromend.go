package linkedlist

func RemoveNthFromEnd(A *Node, B int) *Node {
	count := 1
	head := A
	for head.Next != nil {
		head = head.Next
		count++
	}

	if B >= count {
		head = A.Next
		return head
	}
	curr := A
	k := count - B

	for i := 0; i < (k - 1); i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return A
}

func ReverseFistKElementInGroup(head *Node, k int) *Node {
	curr := head
	count := k
	x := head
	var newNode *Node
	var newHead *Node
	var tail *Node
	for curr != nil && curr.Next != nil {
		newHead, tail, curr = ReverseFirstKElementsNo(curr, count)
		if newNode == nil {
			newNode = newHead
		} else {
			x.Next = newHead
			x = tail
		}
	}

	return newNode
}

func ReverseFirstKElementsNo(head *Node, k int) (newHead *Node, tail *Node, curr *Node) {
	h1 := head
	var h2 *Node
	var tmp *Node
	count := k
	for h1 != nil && count > 0 {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
		count--
	}

	return h2, head, h1
}
