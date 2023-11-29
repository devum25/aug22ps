package greedy

import (
	"container/heap"
	"sort"
)

// You are given an array of events where events[i] = [startDayi, endDayi]. Every event i starts at startDayi and ends at endDayi.
// You can attend an event i at any day d where startTimei <= d <= endTimei. You can only attend one event at any time d.
// Return the maximum number of events you can attend.
// Example 1:
// Input: events = [[1,2],[2,3],[3,4]]

func MaxEvent(events [][]int) int {
	arr := events
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	ans := 0
	day := 0
	index := 0
	n := len(events)

	pq := &intHeap{}
	heap.Init(pq)

	for len(*pq) > 0 || index < n {
		if len(*pq) == 0 {
			day = events[index][0]
		}
		for index < n && events[index][0] == day {
			heap.Push(pq, events[index][1])
			index++
		}
		heap.Pop(pq)
		day++
		ans++
		for len(*pq) > 0 && pq.Top() < day {
			heap.Pop(pq)
		}
	}

	return ans
}

type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *intHeap) Top() int {
	old := *h
	// n := len(old)
	x := old[0]
	return x
}
