package greedy

import (
	"container/heap"
	"math"
)

// Q4. Assign Mice to Holes
// Problem Description
// N Mice and N holes are placed in a straight line. Each hole can accommodate only one mouse.
// The positions of Mice are denoted by array A, and the position of holes is denoted by array B.
// A mouse can stay at his position, move one step right from x to x + 1, or move one step left from x to x − 1. Any of these moves consume 1 minute.
// Assign mice to holes so that the time when the last mouse gets inside a hole is minimized.

// Problem Constraints
// 1 <= N <= 105
// -109 <= A[i], B[i] <= 109

// Input Format
// The first argument is an integer array A.
// The second argument is an integer array B.

// Output Format
// Return an integer denoting the minimum time when the last nouse gets inside the holes.

// Example Input
// Input 1:

//  A = [-4, 2, 3]
//  B = [0, -2, 4]
// Input 2:

//  A = [-2]
//  B = [-6]

// Example Output
// Output 1:
//  2
// Output 2:
//  4
func Mice(A []int, B []int) int {
	m := &minHeap{}
	heap.Init(m)

	for i := 0; i < len(A); i++ {
		heap.Push(m, A[i])
	}

	h := &minHeap{}
	heap.Init(h)

	for i := 0; i < len(B); i++ {
		heap.Push(h, B[i])
	}
	max := 0
	for len(*m) > 0 {
		mice := heap.Pop(m)
		hole := heap.Pop(h)

		temp := int(math.Abs(float64(mice.(int)) - float64(hole.(int))))
		if temp > max {
			max = temp
		}
	}

	return max
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
