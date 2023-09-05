package linkedlist

// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

// You may not alter the values in the list's nodes, only nodes themselves may be changed.

// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
func ReverseKGroup(head *Node, k int) *Node {
	if head == nil || k <= 1 {
		return head
	}

	len := 0

	q := head

	for q != nil {
		q = q.Next
		len++
	}

	if len < k {
		return head
	}

	h1 := head
	var h2 *Node
	var tmp *Node
	count := k
	x := 0
	for h1 != nil && count > 0 {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
		count--
		x++
	}

	if x != k {
		return head
	}

	head.Next = ReverseKGroup(h1, k)

	return h2
}
