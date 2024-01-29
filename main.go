package main

import (
	"github.com/devum25/techbench/dynamicprogramming"
	"github.com/devum25/techbench/graphs"
)

func main() {

	// queue := queue.NewQueue()
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Enqueue(5)
	// stack := queue.ConstructorNew() // i,i,i,i,i,i,i,i,i,i
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)

	// fmt.Print(stack.Top())

	//fmt.Print(queue.MaxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4)) //10,9,9,9,7,11,11

	// fmt.Print(queue.Dequeue())
	// fmt.Print(queue.Dequeue())
	// queue.Enqueue(4)
	// fmt.Print(queue.Dequeue())
	// fmt.Print(queue.Dequeue())
	// fmt.Print(queue.Dequeue())
	// fmt.Print(queue.Dequeue())
	// fmt.Print(math.Pow(34.0515, -3))
	// fmt.Print(binarytree.BuildTree([]int{4, 8, 2, 5, 1, 6, 3, 7}, []int{8, 4, 5, 2, 6, 7, 3, 1}))
	// fmt.Print(binarytree.BuildTree_2([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Enqueue(5)

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
	// head1 := &linkedlist.Node{Val: 1}

	// linkedlist.InsertAtEnd(head1, 2)

	// linkedlist.InsertAtEnd(head1, 3)
	// linkedlist.InsertAtEnd(head1, 4)
	// linkedlist.InsertAtEnd(head1, 5)
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

	// fmt.Print(linkedlist.RotateRight(head1, 10))

	// head = linkedlist.FindMidNodeFirstMid(head)

	// fmt.Print(head1)

	// arr := []int{12, 9, 2, 10, 23, 19, 33, 1, 7}
	// sorting.Mergesort(arr)
	// fmt.Print(arr)

	// dll := queue.NewListNode(2)
	// h := dll.InsertAtStart(1)
	// h.InsertAtEnd(3)
	// h.InsertAtEnd(4)
	// h = h.InsertAtStart(0)

	// fmt.Print(h)

	// deq := queue.NewDeque(1)
	// deq.PushFront(0)
	// deq.PushBack(2)
	// deq.PushBack(3)
	// fmt.Print(deq.RemoveFront())
	// fmt.Print(deq.RemoveBack())

	// }

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
	// tree := binarytree.TreeNode{Val: 1,
	// 	Left: &binarytree.TreeNode{Val: 2,
	// 		Left:  &binarytree.TreeNode{Val: 4},
	// 		Right: &binarytree.TreeNode{Val: 5}},
	// 	Right: &binarytree.TreeNode{Val: 3,
	// 		Left: &binarytree.TreeNode{Val: 6,
	// 			Left:  &binarytree.TreeNode{Val: 9},
	// 			Right: &binarytree.TreeNode{Val: 10}},
	// 		Right: &binarytree.TreeNode{Val: 11}}}

	// x := binarysearchtree.IsHeightBalanced(&tree)

	// fmt.Print(x)
	// node := &binarytree.TreeNode{Val: 8,
	// 	Left:  &binarytree.TreeNode{Val: 9},
	// 	Right: &binarytree.TreeNode{Val: -6, Left: &binarytree.TreeNode{Val: 5}, Right: &binarytree.TreeNode{Val: 9}}}

	// node := &binarytree.TreeNode{Val: 1,
	// 	Left:  &binarytree.TreeNode{Val: 2},
	// 	Right: &binarytree.TreeNode{Val: 3}}
	// Left: &binarytree.TreeNode{Val: 3, Left: &binarytree.TreeNode{Val: 4,
	// 	Left: &binarytree.TreeNode{Val: 5}}}}}

	// node := &binarytree.TreeNode{Val: 1,
	// 	Left: &binarytree.TreeNode{Val: 2}}
	// Left: &binarytree.TreeNode{Val: 3, Left: &binarytree.TreeNode{Val: 4,
	// 	Left: &binarytree.TreeNode{Val: 5}}}}}
	// 	Left: &binarytree.TreeNode{Val: 11, Left: &binarytree.TreeNode{Val: 7}, Right: &binarytree.TreeNode{Val: 2}},
	// },
	// Right: &binarytree.TreeNode{Val: 8, Left: &binarytree.TreeNode{Val: 13},
	// Right: &binarytree.TreeNode{Val: 20, Left: &binarytree.TreeNode{Val: 15}, Right: &binarytree.TreeNode{Val: 7}}}

	// binarytree.MaxPathSum(node)

	// node := &binarytree.TreeNode{Val: 1,
	// 	Left:  &binarytree.TreeNode{Val: 2, Left: &binarytree.TreeNode{Val: 4}, Right: &binarytree.TreeNode{Val: 5}},
	// 	Right: &binarytree.TreeNode{Val: 3, Left: &binarytree.TreeNode{Val: 6}, Right: &binarytree.TreeNode{Val: 7}}}

	// binarytree.MorrisIteration(node)

	// binarytree.DeserializeBt([]int{1, 2, 3, 4, 5, -1, 6, -1, -1, -1, -1, -1, -1})

	// arr := make([][]int, 0)

	// arr = append(arr, []int{1, 2, 3, 4})
	// arr = append(arr, []int{5, 6, 7, 8})
	// arr = append(arr, []int{9, 10, 11, 12})

	// darraygo.SpiralOfRecatange(arr)
	// arrayquestion.MaxAbsoluteSum([]int{2, -5, 1, -4, 3, -2})

	// B := make([][]int, 0)

	// B = append(B, []int{1, 2, 10})
	// B = append(B, []int{2, 3, 20})
	// B = append(B, []int{2, 5, 25})

	// arrayquestion.Flip("010")

	// t := trie.Constructor()

	// t.Insert("FIRE")
	// fmt.Print(t.Search("FIRE"))
	// fmt.Print(t.Search("FIR"))

	// trie.ContactFinder([]int{0, 0, 1, 1}, []string{"hack", "hacker", "hac", "hak"})

	// trie := trie.ConstructorMagicDictionary()

	// trie.BuildDict([]string{"hello", "leetcode"})

	// trie.Search("hhllo")

	// fmt.Print(heaps.MinLargest([]int{5, 7, 8}, 9))
	// // X.Add(-3)
	// // X.Add(-2)
	// // X.Add(-4)
	// // X.Add(0)
	// // X.Add(4)

	// maxHeap := heaps.NewMaxHeap()
	// heaps.SpecialMedian([]int{2147483647, -2147483648, 0})

	// heaps.RunnningMedian([]int{10, 47, 82, 30, 52, 46, 84, 47, 97, 38})
	// heaps.RunnningMedian([]int{1, 2, 5, 4, 3})

	// fmt.Println(heaps.Median([]int{9, 6, 3, 10, 4}))

	// // heaps.HeapSort([]int{8, 3, 7, 6, 1, 5, 10, 4, 9})

	// arr := []int{18, 3, 16, 10, 17, 1, 5, 7, 13, 15}

	// for i := 0; i < len(arr); i++ {
	// 	maxHeap.Insert(arr[i])
	// }

	// fmt.Print(heaps.CreateMaxHeap(arr))

	// fmt.Print(maxHeap.Items)

	// fmt.Print(heaps.MergeRopes([]int{16, 7, 3, 5, 9, 8, 6, 15}))

	// heaps.CallFunc()
	// greedy.MaxNumberOfJobs([]int{4, 4, 8, 15, 6}, []int{9, 5, 15, 16, 7})
	// greedy.BuyTickets(4, 3, []int{2, 2, 2})
	// t := "mjmqqjrmzkvhxlyruonekhhofpzzslupzojfuoztvzmmqvmlhgqxehojfowtrinbatjujaxekbcydldglkbxsqbbnrkhfdnpfbuaktupfftiljwpgglkjqunvithzlzpgikixqeuimmtbiskemplcvljqgvlzvnqxgedxqnznddkiujwhdefziydtquoudzxstpjjitmiimbjfgfjikkjycwgnpdxpeppsturjwkgnifinccvqzwlbmgpdaodzptyrjjkbqmgdrftfbwgimsmjpknuqtijrsnwvtytqqvookinzmkkkrkgwafohflvuedssukjgipgmypakhlckvizmqvycvbxhlljzejcaijqnfgobuhuiahtmxfzoplmmjfxtggwwxliplntkfuxjcnzcqsaagahbbneugiocexcfpszzomumfqpaiydssmihdoewahoswhlnpctjmkyufsvjlrflfiktndubnymenlmpyrhjxfdcq"
	// s :="rjufvjafbxnbgriwgokdgqdqewn"
	//fmt.Println(backtracking.MazeSolver([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
	//fmt.Println(backtracking.SolveNQueensCount(4))

	// backtracking.ValidateSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'6', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}})

	// fmt.Println(backtracking.CountSubsetWithSumK([]int{1, 2, 3, 4, 5, 6}, 4))
	// dfs.AcquireConflictedArea([][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}})
	// graphs.DetectCycle(7, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 4}, {2, 5}, {2, 7}, {4, 6}, {5, 6}})
	dynamicprogramming.TargetSum(2, 6, 7)
	graphs.SelectBatchs(7, []int{1, 6, 7, 2, 9, 4, 5}, [][]int{{1, 2}, {2, 3}, {5, 6}, {5, 7}}, 12)
	// {3, 4}, {1, 5}, {1, 2}, {4, 6}, {2, 3}, {1, 6}, {3, 5}, {0, 3}, {4, 5}, {0, 1}, {0, 5}}
}

// 4,2,5,1,6,7,3
