package linkedlist

func SortDescendingOrder(l1 *MyLinkedList, l2 *MyLinkedList) *MyLinkedList {
	if l2 == nil {
		return reversell(l1)
	} else if l1 == nil {
		return reversell(l2)
	}
	var head *MyLinkedList
	h1 := l1
	h2 := l2
	if h1.Val < h2.Val {
		temp := h1
		h1 = h1.Next
		temp.Next = head
		head = temp
	} else {
		temp := h2
		h2 = h2.Next
		temp.Next = head
		head = temp
	}

	newHead := head

	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			// insert at start
			temp := h1
			h1 = h1.Next
			temp.Next = newHead
			newHead = temp

		} else {
			// insert at start
			temp := h2
			h2 = h2.Next
			temp.Next = newHead
			newHead = temp
		}
	}

	if h1 != nil {
		for h1 != nil {
			temp := h1
			h1 = h1.Next
			temp.Next = newHead
			newHead = temp
		}
	}

	if h2 != nil {
		for h2 != nil {
			temp := h2
			h2 = h2.Next
			temp.Next = newHead
			newHead = temp
		}
	}

	return newHead
}

func reversell(node *MyLinkedList) *MyLinkedList {
	h1 := node
	var h2 *MyLinkedList

	for h1 != nil {
		temp := h1
		h1 = h1.Next
		temp.Next = h2
		h2 = temp
	}

	return h2
}
