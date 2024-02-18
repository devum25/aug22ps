package linkedlist

func Intersection(node1 *MyLinkedList, node2 *MyLinkedList) *MyLinkedList {

	l1 := 0
	l2 := 0

	h1 := node1
	h2 := node2

	for h1 != nil {
		h1 = h1.Next
		l1++
	}

	for h2 != nil {
		h2 = h2.Next
		l2++
	}
	h1 = node1
	if l1 > l2 {
		diff := l1 - l2
		for diff != 0 {
			h1 = h1.Next
			diff--
		}
	}
	h2 = node2
	if l2 > l1 {
		diff := l2 - l1
		for diff != 0 {
			h2 = h2.Next
			diff--
		}
	}

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h2
		}
		h1 = h1.Next
		h2 = h2.Next
	}

	return nil
}
