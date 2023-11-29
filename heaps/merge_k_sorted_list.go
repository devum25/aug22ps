package heaps

import "github.com/devum25/techbench/linkedlist"

// Problem Description
// Given a list containing head pointers of N sorted linked lists.
// Merge these given sorted linked lists and return them as one sorted list.

// Problem Constraints
// 1 <= total number of elements in given linked lists <= 100000
// Input Format
// The first and only argument is a list containing N head pointers.
// Output Format
// Return a pointer to the head of the sorted linked list after merging all the given linked lists.

// Example Input
// Input 1:
//  1 -> 10 -> 20
//  4 -> 11 -> 13
//  3 -> 8 -> 9

// 1 -> 3 -> 4 -> 8 -> 9 -> 10 -> 11 -> 13 -> 20

func CallFunc() {
	x := make([]*linkedlist.Node, 0)
	h := &linkedlist.Node{}
	x = append(x, h)
	MergeKSortedList(x)
}

func MergeKSortedList(lists []*linkedlist.Node) *linkedlist.Node {
	if len(lists) == 0 {
		return nil
	}
	min := NewMinHeap()

	for _, list := range lists {
		for list != nil {
			min.Insert(list.Val)
			list = list.Next
		}
	}

	head := &linkedlist.Node{Val: min.DeleteMin()}
	curr := head

	for len(min.Items) > 0 {
		val := min.DeleteMin()
		node := &linkedlist.Node{Val: val}
		curr.Next = node
		curr = curr.Next
	}

	return head
}
