package linkedlist

func MergeSortedList(head1 *Node, head2 *Node) *Node {

	curr1 := head1
	curr2 := head2

	if curr1 == nil {
		return curr2
	} else if curr2 == nil {
		return curr1
	}

	var head *Node

	if curr1.Val > curr2.Val {
		head = curr2
		curr2 = curr2.Next
	} else {
		head = curr1
		curr1 = curr1.Next
	}

	h := head

	for curr1 != nil && curr2 != nil {

		if curr2.Val < curr1.Val {
			h.Next = curr2
			h = curr2
			curr2 = curr2.Next
		} else {
			h.Next = curr1
			h = curr1
			curr1 = curr1.Next
		}
	}

	if curr1 != nil {
		h.Next = curr1
	} else if curr2 != nil {
		h.Next = curr2
	}

	return head
}
