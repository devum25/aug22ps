package arrayquestion

import "container/heap"

func UnmarkedSumArray(nums []int, queries [][]int) []int64 {
	res := make([]int64, len(queries))
	marked := make(map[int]bool)
	pq := NewMinHeap()
	heap.Init(pq)
	sum := 0
	for i := 0; i < len(nums); i++ {
		heap.Push(pq, MyHeap{Val: nums[i], Idx: i})
		sum += nums[i]
	}

	for i := 0; i < len(queries); i++ {

		idx := queries[i][0]
		k := queries[i][1]

		if !marked[idx] {
			marked[idx] = true
			sum = sum - nums[idx]
		}

		for k > 0 && pq.Len() > 0 {
			temp := heap.Pop(pq).(MyHeap)
			if !marked[temp.Idx] {
				marked[temp.Idx] = true
				sum = sum - temp.Val
				k--
			}
		}

		res[i] = int64(sum)

	}

	return res
}

type MyHeap struct {
	Val int
	Idx int
}

type minHeap []MyHeap

func NewMinHeap() *minHeap {
	return &minHeap{}
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(MyHeap))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Less(i, j int) bool {
	if h[i].Val == h[j].Val {
		return h[i].Idx < h[j].Idx
	}
	return h[i].Val < h[j].Val
}
