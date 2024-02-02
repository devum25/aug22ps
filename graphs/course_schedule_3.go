package graphs

import "container/heap"

// There are n different online courses numbered from 1 to n. You are given an array courses where courses[i] = [durationi, lastDayi] indicate that
// the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.
// You will start on the 1st day and you cannot take two or more courses simultaneously.
// Return the maximum number of courses that you can take.

// Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
// Output: 3

func ScheduleCourse(courses [][]int) int {
	pq := &Heap{}
	heap.Init(pq)

	for i := 0; i < len(courses); i++ {
		heap.Push(pq, CourseSchedule{Duration: courses[i][0], LastDay: courses[i][1]})
	}
	days := 0
	ans := 0
	for pq.Len() > 0 {
		top := heap.Pop(pq).(CourseSchedule)
		if (top.Duration + days) <= top.LastDay {
			days = top.Duration + days
			ans++
		}
	}

	return ans
}

type CourseSchedule struct {
	Duration int
	LastDay  int
}

type Heap []CourseSchedule

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {
	return h[i].Duration < h[j].Duration
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(CourseSchedule))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
