package linkedlist

func Partition(head *Node, x int) *Node {
	if head == nil {
		return nil
	}
	var reliable *Node
	curr := head
	if head.Val < x {
		reliable = head
		curr = curr.Next
	}
	prev := curr
	for curr != nil {
		if curr.Val < x {
			if reliable != nil {
				temp := reliable.Next
				prev.Next = curr.Next
				reliable.Next = curr
				reliable = curr
				curr.Next = temp
				curr = prev.Next
				continue
			} else {
				prev.Next = curr.Next
				curr.Next = head
				head = curr
				curr = prev.Next
				continue
			}
		}
		prev = curr
		curr = curr.Next
	}

	return head
}
