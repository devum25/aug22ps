package linkedlist

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list. Return the linked list sorted as well.
func RemoveDuplicate(head *Node) *Node {
	h := head

	var prev *Node
	for h != nil && h.Next != nil {
		found := false
		for h.Next != nil && h.Val == h.Next.Val {
			h.Next = h.Next.Next
			found = true
		}

		if found {
			if prev != nil {
				prev.Next = h.Next
				h = h.Next
				continue
			} else {
				h = h.Next
				head = h
				continue
			}
		}

		prev = h
		h = h.Next

	}

	return head
}
