package minimumspanningtree

import (
	"container/heap"
	"math"
)

func MinCostConnectPoints(points [][]int) int {
	pq := &Heap{}
	heap.Init(pq)
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			weight := int(math.Abs(float64(points[i][0]-points[j][0]))) + int(math.Abs(float64(points[i][1]-points[j][1])))
			heap.Push(pq, MinHeap{Node1: i, Node2: j, Weight: weight})
		}
	}

	parent := make([]int, n)
	height := make([]int, n)
	for i := range parent {
		parent[i] = i
		height[i] = 0
	}

	cost := 0
	for pq.Len() > 0 {
		top := heap.Pop(pq).(MinHeap)

		if union(top.Node1, top.Node2, parent, height) {
			cost += top.Weight
		}
	}

	return cost
}

func union(u, v int, parent, height []int) bool { // takes two vertices and if it is possible to connect them then does it

	ur := pathCompression(u, parent)
	vr := pathCompression(v, parent)

	if ur != vr {
		// parent[ur] = uv
		hur := height[ur]
		hvr := height[vr]

		if hur > hvr {
			parent[vr] = ur
		} else if hvr > hur {
			parent[ur] = vr
		} else {
			parent[ur] = vr
			height[vr]++
		}

		return true
	}
	return false
}

func findRoot(u int, parent []int) int { // find root of vertices

	i := u
	for parent[i] != i {
		i = parent[i]
	}

	return i
}

func pathCompression(u int, parent []int) int {
	if parent[u] == u {
		return u
	}
	parent[u] = pathCompression(parent[u], parent)

	return parent[u]
}

type MinHeap struct {
	Node1  int
	Node2  int
	Weight int
}

type Heap []MinHeap

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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
