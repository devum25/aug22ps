package linkedlist

// NOT CLEAN SOLUTION
func SwapTwoAdjacent(head *Node) *Node {
	curr := head
	var newNode *Node
	var h1 *Node
	var newHead *Node
	for curr != nil {
		newNode, curr = swap(curr)
		if h1 == nil {
			h1 = newNode
			newHead = newNode
		} else {
			h1.Next.Next = newNode
			h1 = newNode
		}

	}
	return newHead
}

func swap(curr *Node) (*Node, *Node) {
	if curr.Next != nil {
		tmp := curr.Next
		curr.Next = tmp.Next
		tmp.Next = curr
		return tmp, curr.Next
	}
	return curr, nil
}

// can i say this problem as reverse 2 node in group?
func SwapTwoAdjacentCleanApproach(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	h1 := head
	var h2 *Node
	k := 0
	for h1 != nil && k < 2 {
		tmp := h1
		h1 = tmp.Next
		tmp.Next = h2
		h2 = tmp
		k++
	}

	head.Next = SwapTwoAdjacentCleanApproach(h1)
	return h2
}
