package arrayquestion

import (
	"container/heap"
)

// 786. K-th Smallest Prime Fraction
// Medium
// Topics
// Companies
// You are given a sorted integer array arr containing 1 and prime numbers, where all the integers of arr are unique. You are also given an integer k.

// For every i and j where 0 <= i < j < arr.length, we consider the fraction arr[i] / arr[j].

// Return the kth smallest fraction considered. Return your answer as an array of integers of size 2, where answer[0] == arr[i] and answer[1] == arr[j].

// Example 1:

// Input: arr = [1,2,3,5], k = 3
// Output: [2,5]
// Explanation: The fractions to be considered in sorted order are:
// 1/5, 1/3, 2/5, 1/2, 3/5, and 2/3.
// The third fraction is 2/5.
func KthSmallestPrimeFraction(arr []int, k int) []int {
	pq := &MinHeap{}
	heap.Init(pq)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(pq, Heap{Item: float64(float64(arr[i]) / float64(arr[j])), I: i, J: j})
		}
	}
	ans := make([]int, 0)
	for k > 0 {
		top := heap.Pop(pq).(Heap)
		k--
		if k == 0 {
			ans = append(ans, arr[top.I])
			ans = append(ans, arr[top.J])
		}
	}

	return ans
}

type Heap struct {
	Item float64
	I    int
	J    int
}

type MinHeap []Heap

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].Item < m[j].Item
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(Heap))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Optimal solution
// https://www.youtube.com/watch?v=ZwbSRpPOVHg
func kthSmallestPrimeFraction(arr []int, k int) []int {
	start := float64(float64(arr[0]) / float64(arr[len(arr)-1]))
	var end float64
	end = 1.0

	for start < float64(end) {
		mid := float64((start + float64(end)) / 2)

		count, nemo, deno := getCount(arr, mid)

		if count > k {
			end = mid
		} else if count < k {
			start = mid
		} else {
			return []int{nemo, deno}
		}
	}

	return nil
}

func getCount(arr []int, mid float64) (int, int, int) {
	nemo := arr[0]
	deno := arr[len(arr)-1]
	i := 0
	count := 0
	for j := 1; j < len(arr); j++ {
		for float64(arr[i]) <= float64(arr[j])*mid {
			i++
		}
		count += i
		if i > 0 && arr[i-1]*deno > arr[j]*nemo {
			nemo = arr[i-1]
			deno = arr[j]
		}
	}

	return count, nemo, deno
}
