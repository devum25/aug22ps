package linkedlist

import "fmt"

func MergeSortLinkedList(head *Node) *Node {
	return mergelinkedlist(head)
}

func mergelinkedlist(head *Node) *Node {

	if head == nil || head.Next == nil {
		return head
	}

	mid := findMidOfLinkedList(head)
	h2 := mid.Next
	mid.Next = nil
	h1 := mergelinkedlist(head)
	h2 = mergelinkedlist(h2)
	return mergeTwoSortedLinkedList(h1, h2)
}

func mergeTwoSortedLinkedList(h1 *Node, h2 *Node) *Node {
	curr1 := h1
	curr2 := h2

	if curr1 == nil {
		return curr2
	} else if curr2 == nil {
		return curr1
	}

	var head *Node
	if curr1.Val < curr2.Val {
		head = curr1
		curr1 = curr1.Next
	} else {
		head = curr2
		curr2 = curr2.Next
	}

	tail := head

	for curr1 != nil && curr2 != nil {

		if curr2.Val < curr1.Val {
			tail.Next = curr2
			tail = curr2
			curr2 = curr2.Next
		} else {
			tail.Next = curr1
			tail = curr1
			curr1 = curr1.Next
		}

	}

	if tail == nil {
		fmt.Print("a ...any")
	}

	if curr1 != nil {
		tail.Next = curr1
	} else if curr2 != nil {
		tail.Next = curr2
	}

	return head
}

func findMidOfLinkedList(head *Node) *Node {
	slow := head
	fast := head
	prev := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		return prev
	}

	return slow
}
