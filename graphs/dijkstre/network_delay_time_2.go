package dijkstre

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	adjList := make(map[int][][2]int)

	for i := 0; i < len(times); i++ {
		adjList[times[i][0]] = append(adjList[times[i][0]], [2]int{times[i][1], times[i][2]})
	}

	lst := make([]int, n+1)

	for i := 0; i < len(lst); i++ {
		lst[i] = math.MaxInt
	}

	lst[0] = 0
	lst[k] = 0

	pq := &Heap1{}
	heap.Init(pq)

	heap.Push(pq, MinHeap1{Node: k, Dist: 0})
	min := math.MinInt
	for pq.Len() > 0 {
		temp := heap.Pop(pq).(MinHeap1)

		if temp.Dist == lst[temp.Node] {
			for _, v := range adjList[temp.Node] {
				if v[1]+temp.Dist < lst[v[0]] {
					heap.Push(pq, MinHeap1{Node: v[0], Dist: v[1] + temp.Dist})
					lst[v[0]] = temp.Dist + v[1]
				}
			}
		}
	}

	for i := 0; i < len(lst); i++ {
		if lst[i] > min {
			min = lst[i]
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min

}

type MinHeap1 struct {
	Node int
	Dist int
}

type Heap1 []MinHeap1

func (h Heap1) Len() int           { return len(h) }
func (h Heap1) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h Heap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap1) Push(x interface{}) {
	*h = append(*h, x.(MinHeap1))
}

func (h *Heap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *Heap1) Top() MinHeap1 {
	old := *h
	x := old[0]
	return x
}
