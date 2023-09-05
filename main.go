package main

import (
	"fmt"

	"github.com/devum25/techbench/linkedlist"
)

// [4,6,7,4,6,9,7,2,3,6] //9
func main() {
	head := &linkedlist.Node{Val: 1}

	for i := 2; i <= 15; i++ {
		linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
	}

	fmt.Print(head)

	// head = linkedlist.DeleteAtStart(head)
	// head = linkedlist.DeleteAtEnd(head)
	head = linkedlist.ReverseInKGroupUsinRecursion(head, 4)

	fmt.Print(head)

}
