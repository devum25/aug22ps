package linkedlist

func ReverseFirstKElemets(head *MyLinkedList, k int) *MyLinkedList {
	h1 := head
	var h2 *MyLinkedList
	for i := 0; i < k; i++ {
		temp := h1
		h1 = h1.Next
		temp.Next = h2
		h2 = temp
	}

	head.Next = h1
	return h2
}
