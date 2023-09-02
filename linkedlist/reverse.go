package linkedlist

func Reverse(head *Node) *Node {
	h1 := head
	var h2 *Node
	var tmp *Node

	for h1 != nil {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
	}

	return h2
}

func ReverseFirstKElements(head *Node, k int) *Node {
	var h1 *Node
	var h2 *Node
	var tmp *Node
	var firstNode *Node
	h1 = head

	if k == 0 || k == 1 {
		return head
	}

	for i := 0; i < k; i++ {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
		if i == 0 {
			firstNode = h2
		}
	}

	firstNode.Next = h1

	return h2
}

func ReverseEverySublistofKElements(head *Node, k int) *Node {
	var h1 *Node
	var newNode *Node
	var tmp *Node
	h1 = head
	count := k
	first := true
	for h1 != nil {
		count = k
		var h2 *Node
		for h1 != nil && count > 0 {
			tmp = h1
			h1 = h1.Next
			tmp.Next = h2
			h2 = tmp
			count--
		}

		head.Next = h1
		head = h1
		if first {
			newNode = h2
			first = false
		}
	}
	return newNode
}
