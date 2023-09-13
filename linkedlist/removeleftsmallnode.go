package linkedlist

// 5->2->13->3->8->NIL
// 13->8->NIL

// reverse and arrange in ascending order then

func RemoveNodes(head *Node) *Node {
	curr := reverse(head)
	currMax := curr.Val
	h := curr
	prev := curr
	curr = curr.Next
	for curr != nil {
		if curr.Val > currMax {
			currMax = curr.Val
			prev.Next = curr
			prev = curr
		}

		curr = curr.Next
	}

	// if prev.Next != nil {
	prev.Next = nil
	// }
	return reverse(h)
}

func reverse(head *Node) *Node {
	h1 := head
	var h2 *Node
	var tmp *Node

	for h1 != nil {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
	}

	return h2
}
