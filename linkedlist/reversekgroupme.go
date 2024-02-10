package linkedlist

func ReverseKGroup1(head *MyLinkedList, k int) *MyLinkedList {
	count := 0
	curr := head

	for curr != nil {
		curr = curr.Next
		count++
	}
	h := reverse1(head, k, count)
	return h
}

func reverse1(head *MyLinkedList, x int, count int) *MyLinkedList {
	if head == nil || count < x {
		return head
	}
	k := x
	h1 := head
	var h2 *MyLinkedList
	for h1 != nil && k > 0 {
		tmp := h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
		k--
	}
	tail := head
	// if count == total {
	head = h2
	// }
	count = count - x
	tail.Next = reverse1(h1, x, count)

	return head
}
