package main

import (
	"fmt"

	"github.com/devum25/techbench/linkedlist"
)

// [4,6,7,4,6,9,7,2,3,6] //9
func main() {

	// for i := 2; i <= 9; i++ {
	// 	linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
	// }
	//67 -> 27 -> 64 -> 10 -> 4 -> 85
	// fmt.Print(head)
	head1 := &linkedlist.Node{Val: 2}

	linkedlist.InsertAtEnd(head1, 1)
	linkedlist.InsertAtEnd(head1, 2)
	linkedlist.InsertAtEnd(head1, 1)
	linkedlist.InsertAtEnd(head1, 2)
	linkedlist.InsertAtEnd(head1, 2)
	linkedlist.InsertAtEnd(head1, 1)
	linkedlist.InsertAtEnd(head1, 3)
	linkedlist.InsertAtEnd(head1, 2)
	linkedlist.InsertAtEnd(head1, 2)
	// linkedlist.InsertAtEnd(head2, 4)
	// linkedlist.InsertAtEnd(head2, 85)
	// linkedlist.InsertAtEnd(head1, 33)
	// linkedlist.InsertAtEnd(head1, 1)
	// linkedlist.InsertAtEnd(head1, 7)
	// head = linkedlist.DeleteAtEnd(head)

	// head2 := &linkedlist.Node{Val: 2}

	// linkedlist.InsertAtEnd(head2, 6)
	// linkedlist.InsertAtEnd(head2, 7)
	// linkedlist.InsertAtEnd(head2, 11)
	// linkedlist.InsertAtEnd(head2, 12)
	// linkedlist.InsertAtEnd(head2, 15)

	fmt.Print(linkedlist.LongestPalindrome(head1))

	// head = linkedlist.FindMidNodeFirstMid(head)

	fmt.Print(head1)

	// arr := []int{12, 9, 2, 10, 23, 19, 33, 1, 7}
	// sorting.Mergesort(arr)
	// fmt.Print(arr)

}
