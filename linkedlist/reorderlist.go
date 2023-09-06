package linkedlist

import "fmt"

func ReorderList(head *Node) *Node {
	h1 := head
	h2 := reverseList(h1)
	fmt.Print(h2)
	return h2
}

func reverseList(head *Node) *Node {
	var h2 *Node
	h1 := head
	tmp := head
	for h1 != nil {
		tmp = h1
		h1 = h1.Next
		tmp.Next = h2
		h2 = tmp
	}

	return h2
}
