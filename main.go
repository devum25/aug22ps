package main

import (
	"fmt"

	"github.com/devum25/techbench/linkedlist"
	"github.com/devum25/techbench/queue"
)

func main() {

	queue := queue.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(5)

	fmt.Print(queue.Dequeue())
	fmt.Print(queue.Dequeue())
	queue.Enqueue(4)
	fmt.Print(queue.Dequeue())
	fmt.Print(queue.Dequeue())
	fmt.Print(queue.Dequeue())
	fmt.Print(queue.Dequeue())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(5)

	// lru := linkedlist.NewLRUCache(2)
	// lru.Get(2)
	// lru.Put(2, 6)
	// // fmt.Print(lru.Get(1))
	// lru.Put(1, 1)
	// lru.Put(2, 3)
	// lru.Put(4, 1)
	// fmt.Print(lru.Get(1))
	// fmt.Print(lru.Get(2))
	// lru.Get(4)
	// lru.Get(2)
	// lru.Put(1, 1)
	// for i := 2; i <= 9; i++ {
	// 	linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
	// }
	//67 -> 27 -> 64 -> 10 -> 4 -> 85
	// fmt.Print(head)
	head1 := &linkedlist.Node{Val: 1}

	linkedlist.InsertAtEnd(head1, 2)

	linkedlist.InsertAtEnd(head1, 3)
	linkedlist.InsertAtEnd(head1, 4)
	linkedlist.InsertAtEnd(head1, 5)
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

	fmt.Print(linkedlist.RotateRight(head1, 10))

	// head = linkedlist.FindMidNodeFirstMid(head)

	fmt.Print(head1)

	// arr := []int{12, 9, 2, 10, 23, 19, 33, 1, 7}
	// sorting.Mergesort(arr)
	// fmt.Print(arr)

}

// 8,222,222,-1,222,1111111,-1,1
// 59,59,97,58,58,97,-1,-23,84,97
// func main() {
// 	// s := stacks.ConvertToPostFix("3 + 10 * ( 3 - 4 / 2 ) + 3")
// 	// s := stacks.ConvertToPostFix("13+10*(3-4/2)+3")
// 	// x := []int{0, 2, 0}
// 	// fmt.Print(stacks.SmallestSubsequenceWithLetter("wuynymkihfdcbabefiiymnoyyytywzy", 16, 'y', 4))
// 	// fmt.Print(s)
// 	// matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
// 	// fmt.Print(binarysearch.SearchMatrix(matrix, 3)) // 230412
// 	// matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
// 	// matrix := [][]int{{1, 1}}
// 	fmt.Print(binarysearch.MaximumCount([]int{-2, -1, -1, 0, 0, 0}))

// 	// arr := []int{5, 2, 4, 7, 1, 1}
// 	// for i := 0; i < len(arr); i++ {
// 	// 	s.Push(arr[i])
// 	// }

// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Print(s.GetMin())
// 	// 	s.Pop()
// 	// }
// }

//abefiimnoyytywzy
//uynyabefiiymnoyy
