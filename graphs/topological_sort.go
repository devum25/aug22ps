package graphs

import "container/heap"

func TopologicalSort(A int, B [][]int) []int {
	adjList := make(map[int][]int)

	for _, v := range B {
		adjList[v[0]] = append(adjList[v[0]], v[1])
	}

	return toposort(A, adjList)
}

func toposort(n int, adjList map[int][]int) []int {
	indegree := make(map[int]int)
	pq := &IntHeap{}
	heap.Init(pq)
	order := make([]int, 0)
	for i := 1; i <= n; i++ {
		indegree[n] = 0
	}

	for k := range adjList {
		for _, v := range adjList[k] {
			indegree[v]++
		}
	}

	visited := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 && !visited[i] {
			visited[i] = true
			order = append(order, i)
			pq.Push(i)
		}
		for pq.Len() > 0 {
			temp := pq.Pop().(int)
			// queue = queue[1:]

			for _, v := range adjList[temp] {
				indegree[v]--
				if indegree[v] == 0 {
					order = append(order, v)
					visited[v] = true
					pq.Push(v)
				}
			}
		}
	}

	for _, v := range indegree {
		if v > 0 {
			return []int{}
		}
	}

	return order

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() int {
	old := *h
	// n := len(old)
	x := old[0]
	return x
}
