package main

import (
	"fmt"

	"github.com/devum25/techbench/linkedlist"
)

// [4,6,7,4,6,9,7,2,3,6] //9
func main() {
	head1 := &linkedlist.Node{Val: 3}

	// for i := 2; i <= 3; i++ {
	// 	linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
	// }

	// fmt.Print(head)

	linkedlist.InsertAtEnd(head1, 8)
	linkedlist.InsertAtEnd(head1, 10)
	linkedlist.InsertAtEnd(head1, 14)
	linkedlist.InsertAtEnd(head1, 20)
	// head = linkedlist.DeleteAtEnd(head)

	head2 := &linkedlist.Node{Val: 2}

	linkedlist.InsertAtEnd(head2, 6)
	linkedlist.InsertAtEnd(head2, 7)
	linkedlist.InsertAtEnd(head2, 11)
	linkedlist.InsertAtEnd(head2, 12)
	linkedlist.InsertAtEnd(head2, 15)

	head1 = linkedlist.MergeSortedList(head1, head2)

	fmt.Print(head1)

}
