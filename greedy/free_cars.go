package greedy

// Q3. Free Cars
// Unsolved
// feature icon
// Get your doubts resolved blazing fast with Chat GPT Help
// Check Chat GPT
// feature icon
// Using hints is now penalty free
// Use Hint
// Problem Description
// Given two arrays, A and B of size N. A[i] represents the time by which you can buy the ith car without paying any money.

// B[i] represents the profit you can earn by buying the ith car. It takes 1 minute to buy a car, so you can only buy the ith car when the current time <= A[i] - 1.

// Your task is to find the maximum profit one can earn by buying cars considering that you can only buy one car at a time.

// NOTE:

// You can start buying from time = 0.
// Return your answer modulo 109 + 7.

// Problem Constraints
// 1 <= N <= 105
// 1 <= A[i] <= 109
// 0 <= B[i] <= 109

// Input Format
// The first argument is an integer array A represents the deadline for buying the cars.
// The second argument is an integer array B represents the profit obtained after buying the cars.

// Output Format
// Return an integer denoting the maximum profit you can earn.

// Example Input
// Input 1:

//  A = [1, 3, 2, 3, 3]
//  B = [5, 6, 1, 3, 9]
// Input 2:

//  A = [3, 8, 7, 5]
//  B = [3, 1, 7, 19]

// Example Output
// Output 1:

//  20
// Output 2:

//  30

// Example Explanation
// Explanation 1:

//  At time 0, buy car with profit 5.
//  At time 1, buy car with profit 6.
//  At time 2, buy car with profit 9.
//  At time = 3 or after , you can't buy any car, as there is no car with deadline >= 4.
//  So, total profit that one can earn is 20.

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer
 */
import (
	"container/heap"
	"sort"
)

// func FreeCars(A []int, B []int) int {
// 	maxTime := 0
// 	ans := 0

// 	for i := 0; i < len(A); i++ {
// 		if A[i] > maxTime {
// 			maxTime = A[i]
// 		}
// 	}

// 	arr := make([][]int, len(B))

// 	for i := 0; i < len(B); i++ {
// 		temp := make([]int, 2)
// 		temp[0] = i
// 		temp[1] = B[i]
// 		arr[i] = temp
// 	}

// 	sort.Slice(arr, func(i, j int) bool {
// 		return arr[i][1] > arr[j][1]
// 	})

// 	res := make([]int, maxTime)

// 	for i := 0; i < len(arr); i++ {
// 		t := arr[i][0]
// 		idx := A[t] - 1
// 		if idx < len(res) && res[idx] == 0 {
// 			res[idx] = arr[i][1]
// 		} else {
// 			j := idx - 1
// 			for j >= 0 {
// 				if j < len(res) && res[j] == 0 {
// 					res[j] = arr[i][1]
// 					break
// 				}
// 				j--
// 			}
// 		}
// 	}

// 	for i := 0; i < len(res); i++ {
// 		ans = ans + res[i]
// 	}

// 	return int(math.Mod(float64(ans), math.Pow(10, 9)+7))
// }

func FreeCars(A []int, B []int) int {

	res := make([][]int, 0)

	for i := 0; i < len(A); i++ {
		res = append(res, []int{A[i], B[i]})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})

	ans := 0
	pq := &IntHeap{}
	heap.Init(pq)
	n := len(res)
	t := 0
	i := 0
	for i < n {
		// if we can buy the ith car
		if res[i][0] > t {
			heap.Push(pq, res[i][1])
			t += 1
		} else {
			min_profit := heap.Pop(pq).(int)
			//reomve the already buy car giving minimum profit
			if min_profit < res[i][1] {
				heap.Push(pq, res[i][1])
			} else {
				heap.Push(pq, min_profit)
			}
		}
		i += 1
	}

	for pq.Len() > 0 {
		ans += heap.Pop(pq).(int)
	}

	return ans
}

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer
 */

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

func solve(A []int, B []int) int {
	n := len(A)
	var v [][]int
	for i := 0; i < n; i++ {
		v = append(v, []int{A[i], B[i]})
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i][0] < v[j][0]
	})

	curTime := 0
	i := 0
	pq := &IntHeap{}
	heap.Init(pq)

	for i < n {
		// if we can buy the ith car
		if v[i][0] > curTime {
			heap.Push(pq, v[i][1])
			curTime += 1
		} else {
			min_profit := heap.Pop(pq).(int)
			//reomve the already buy car giving minimum profit
			if min_profit < v[i][1] {
				heap.Push(pq, v[i][1])
			} else {
				heap.Push(pq, min_profit)
			}
		}
		i += 1
	}
	var ans int = 0
	var MOD int = 1000000007
	for pq.Len() > 0 {
		ans += heap.Pop(pq).(int)
		ans %= MOD
	}
	return int(ans)
}
