package heaps

import "fmt"

func KSmallerElements(arr []int, k int) []int {

	x := CreateMaxHeap(arr[:k])
	fmt.Print(arr)
	heap := NewMaxHeap()
	heap.Items = x
	heap.size = k - 1

	for i := k; i < len(arr); i++ {
		if arr[i] < arr[0] {
			heap.DeleteMax()
			heap.Insert(arr[i])
		}
	}

	res := make([]int, k)
	i := 0
	for len(heap.Items) > 0 {
		res[i] = heap.DeleteMax()
		i++
	}

	return res
}
