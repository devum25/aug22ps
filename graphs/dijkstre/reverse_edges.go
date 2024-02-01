package dijkstre

import (
	"container/heap"
	"math"
)

// Given a directed graph with A nodes and M edges.
// Find the minimum number of edges that needs to be reversed in order to reach node A from node 1.
func ReverseEdges(A int, B [][]int) int {
	adjList := make(map[int][][2]int)

	for _, v := range B {
		adjList[v[0]] = append(adjList[v[0]], [2]int{v[1], 0})
		adjList[v[1]] = append(adjList[v[1]], [2]int{v[0], 1})
	}

	return reverse(adjList, A)
}

func reverse(adjList map[int][][2]int, n int) int {
	pq := &Heap{}
	heap.Init(pq)

	lst := make([]int, n+1)

	for i := 2; i <= n; i++ {
		lst[i] = math.MaxInt
	}

	pq.Push(MinHeap{Node: 1, Distance: 0})

	for pq.Len() > 0 {
		temp := pq.Pop().(MinHeap)

		if temp.Distance == lst[temp.Node] {
			for _, v := range adjList[temp.Node] {
				if v[1] < lst[v[0]] {
					pq.Push(MinHeap{Node: v[0], Distance: v[1] + temp.Distance})
					lst[v[0]] = v[1] + temp.Distance
				}
			}
		}
	}

	return lst[n]
}
