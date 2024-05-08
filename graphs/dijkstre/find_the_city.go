package dijkstre

import (
	"container/heap"
	"math"
)

func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	pq := &Heap{}
	heap.Init(pq)

	heap.Push(pq, MinHeap{Node: edges[0][0], Distance: 0})

	adjList := make(map[int][][2]int)

	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], [2]int{edges[i][1], edges[i][2]})
	}

	lst := make([]int, n+1)

	cities := make([]int, n+1)

	for i := 0; i < n; i++ {
		lst[i] = math.MaxInt
	}
	lst[0] = 0
	lst[edges[0][0]] = 0
	min := math.MaxInt
	for pq.Len() > 0 {
		top := heap.Pop(pq).(MinHeap)

		neighbours := adjList[top.Node]
		d := top.Distance
		val := 0
		if d == lst[top.Node] {
			for _, v := range neighbours {
				if (v[1]+d) < lst[v[0]] && (v[1]+d) <= distanceThreshold {
					lst[v[0]] = v[1] + d
					heap.Push(pq, MinHeap{Node: v[0], Distance: v[1] + d})
					t := cities[top.Node]
					cities[top.Node] = t + 1
					val++
				}
			}
			if val < min {
				min = val
			}
		}
	}
	ans := 0
	for i := 0; i < len(cities); i++ {
		if cities[i] == min {
			ans = i
		}
	}

	return ans
}

// type MinHeap struct {
// 	Node int
// 	Dist int
// }

// type Heap []MinHeap

// func (h Heap) Len() int           { return len(h) }
// func (h Heap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
// func (h Heap) Swap(i, j int)      { h[i].Dist, h[j].Dist = h[j].Dist, h[i].Dist }

// func (h *Heap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func (h *Heap) Push(x interface{}) {
// 	*h = append(*h, x.(MinHeap))
// }
