package main

import (
	"fmt"

	"github.com/devum25/techbench/linkedlist"
)

// [4,6,7,4,6,9,7,2,3,6] //9
func main() {
	head := &linkedlist.Node{Val: 0}

	for i := 1; i <= 5; i++ {
		linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
	}

	fmt.Print(head)

	// head = linkedlist.DeleteAtStart(head)
	// head = linkedlist.DeleteAtEnd(head)
	head = linkedlist.ReverseEverySublistofKElements(head, 3)

	fmt.Print(head)

}
