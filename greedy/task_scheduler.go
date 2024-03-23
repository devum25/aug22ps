package greedy

import "container/heap"

func LeastInterval(tasks []byte, n int) int {
	hash := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		hash[tasks[i]]++
	}
	pq := NewMaxHeap()
	heap.Init(pq)

	for k, v := range hash {
		heap.Push(pq, MyHeap{Item: k, Count: v})
	}
	arr := make([]byte, 0)

	start := 0
	for pq.Len() > 0 {
		temp := pq.Pop().(MyHeap)

		c := temp.Count

		if len(arr) == 0 {
			arr = append(arr, temp.Item)
			c--
			start = 1
			for c > 0 {
				for i := 1; i <= n; i++ {
					arr = append(arr, '-')
				}
				arr = append(arr, temp.Item)
				c--
			}
		} else {
			if arr[start] == '-' {
				arr[start] = temp.Item
				c--
				start = start + 1
				for c > 0 {
					k := start
					for i := 1; i <= n; i++ {
						if k >= len(arr) {
							arr = append(arr, '-')
						}
						k++
					}
					if k < len(arr) {
						arr[k] = temp.Item
						start = k + 1
					} else {
						arr = append(arr, temp.Item)
						arr = append(arr, '-')
						start = len(arr) - 1
					}
					c--
				}
			}
		}
	}

	return len(arr)
}

type MyHeap struct {
	Item  byte
	Count int
}

type maxHeap1 []MyHeap

func NewMaxHeap() *maxHeap1 {
	return &maxHeap1{}
}

func (m *maxHeap1) Push(x interface{}) {
	*m = append(*m, x.(MyHeap))
}
func (m *maxHeap1) Pop() interface{} {
	h := *m
	n := len(h)
	val := h[n-1]
	h = h[:n-1]
	*m = h
	return val
}

func (m maxHeap1) Len() int {
	return len(m)
}
func (m maxHeap1) Less(i, j int) bool {
	return m[i].Count < m[j].Count
}
func (m maxHeap1) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
