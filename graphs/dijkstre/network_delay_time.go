package dijkstre

import (
	"container/heap"
	"math"
)

// You are given a network of n nodes, labeled from 1 to n. You are also given times,
//  a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node,
//  vi is the target node, and wi is the time it takes for a signal to travel from source to target.
// We will send a signal from a given node k.
// Return the minimum time it takes for all the n nodes to receive the signal.
// If it is impossible for all the n nodes to receive the signal, return -1.

func NetworkDelayTime(times [][]int, n int, k int) int {

	adjList := make(map[int][][2]int)

	for _, v := range times {
		adjList[v[0]] = append(adjList[v[0]], [2]int{v[1], v[2]})
	}

	return dijkstre(adjList, k, n)
}

func dijkstre(adjList map[int][][2]int, source int, nodes int) int {
	pq := &Heap{}
	heap.Init(pq)
	min := math.MinInt
	pq.Push(MinHeap{Node: source, Distance: 0})
	lst := make([]int, nodes+1)
	for i := 0; i < len(lst); i++ {
		lst[i] = math.MaxInt
	}
	lst[0] = 0
	lst[source] = 0
	for pq.Len() > 0 {
		temp := pq.Pop().(MinHeap)

		if temp.Distance == lst[temp.Node] {
			for _, v := range adjList[temp.Node] {
				if v[1]+temp.Distance < lst[v[0]] {
					pq.Push(MinHeap{Node: v[0], Distance: v[1] + temp.Distance})
					lst[v[0]] = temp.Distance + v[1]
				}
			}
		}
	}

	for i := 1; i < len(lst); i++ {
		if lst[i] > min {
			min = lst[i]
		}
	}

	if min == math.MaxInt {
		return -1
	}

	return min
}

type MinHeap struct {
	Node     int
	Distance int
}

type Heap []MinHeap

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h Heap) Swap(i, j int)      { h[i].Distance, h[j].Distance = h[j].Distance, h[i].Distance }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(MinHeap))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Top() MinHeap {
	old := *h
	// n := len(old)
	x := old[0]
	return x
}
