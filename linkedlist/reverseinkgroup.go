package linkedlist

func ReverseInKGroupUsinRecursion(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}
	h1 := head
	var h2 *Node
	var tmp *Node
	count := k
	for h1 != nil && count > 0 {
		tmp = h1
		h1 = tmp.Next
		tmp.Next = h2
		h2 = tmp
		count--
	}

	head.Next = ReverseInKGroupUsinRecursion(h1, k)

	return h2
}
