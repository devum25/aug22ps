package linkedlist

func ReverseBetween(head *Node, B int, C int) *Node {
	k := C - B + 1
	if k <= 1 || head == nil {
		return head
	}

	if B == 1 {
		newHead := ReverseFirstKElements(head, k)
		return newHead
	}

	curr := head
	x := 1
	prev := head
	for curr != nil {
		prev = curr
		curr = curr.Next
		x++
		if x == B {
			newHead := ReverseFirstKElements(curr, k)
			prev.Next = newHead
			return head
		}
	}

	return nil
}
