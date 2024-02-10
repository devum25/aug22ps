package linkedlist

func ReverseBetween3(head *MyLinkedList, left int, right int) *MyLinkedList {
	h1 := head
	count := 1
	curr := head

	for curr != nil && count < left {
		h1 = curr
		curr = curr.Next
		count++
	}

	h1.Next = reverse3(curr, right-left+1)
	return head
}

func reverse3(head *MyLinkedList, k int) *MyLinkedList {
	h1 := head
	var h2 *MyLinkedList
	for h1 != nil && k > 0 {
		temp := h1
		h1 = h1.Next
		temp.Next = h2
		h2 = temp
		k--
	}
	head.Next = h1
	return h2
}
