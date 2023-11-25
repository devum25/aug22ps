package heaps

import "fmt"

// TC : O(NlogN)
// SC : O(1)
func HeapSort(arr []int) {
	heap := NewHeapSortNode()
	ConvertToMaxHeap(arr)
	fmt.Print(arr)
	heap.Item = arr
	heap.Size = len(arr) - 1

	i := 0

	for i < len(heap.Item) {
		heap.ExtractMax()
		i++
	}
	return
}

type HeapSortNode struct {
	Item []int
	Size int
}

func NewHeapSortNode() *HeapSortNode {
	return &HeapSortNode{Item: make([]int, 0)}
}

func (h *HeapSortNode) ExtractMax() {
	h.Item[0], h.Item[h.Size] = h.Item[h.Size], h.Item[0]
	h.Size--

	i := 0

	for i <= h.Size {
		max := i
		l := 2*i + 1
		r := 2*i + 2

		if l <= h.Size && h.Item[l] > h.Item[max] {
			max = l
		}
		if r <= h.Size && h.Item[r] > h.Item[max] {
			max = r
		}

		if max == i {
			break
		}

		h.Item[max], h.Item[i] = h.Item[i], h.Item[max]
		i = max
	}
	return
}

func ConvertToMaxHeap(arr []int) {
	leaf := (len(arr) + 1) / 2
	s := leaf - 1
	for s >= 0 {
		i := s
		for i < len(arr) {
			max := i
			l := 2*i + 1
			r := 2*i + 2
			if l < len(arr) && arr[l] > arr[max] {
				max = l
			}
			if r < len(arr) && arr[r] > arr[max] {
				max = r
			}
			if i == max {
				break
			}
			arr[max], arr[i] = arr[i], arr[max]
			i = max
		}
		s--
	}
	return
}
