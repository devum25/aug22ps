package linkedlist

func RotateRight(head *Node, k int) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	h := head
	l := getLength(h)
	x := l - k%l
	if k == 0 || k == l || x == l {
		return head
	}

	if x == 0 {
		return head
	}

	count := 0
	var pivot *Node
	for h != nil {
		count++
		if count == x {
			pivot = h.Next
			h.Next = nil
			break
		}
		h = h.Next
	}

	h1 := pivot

	for h1 != nil && h1.Next != nil {
		h1 = h1.Next
	}

	if h1 != nil {
		h1.Next = head
	}

	return pivot
}

func getLength(head *Node) int {
	count := 0
	h := head

	for h != nil {
		h = h.Next
		count++
	}

	return count
}
