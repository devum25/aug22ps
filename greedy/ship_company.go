package greedy

import "container/heap"

// Q3. The ship company
// Problem Description
// The local ship renting service has a special rate plan:
// It is up to a passenger to choose a ship.
// If the chosen ship has X (X > 0) vacant places at the given moment, then the ticket for such a ship costs X.
// The passengers buy tickets in turn, the first person in the queue goes first, then the second one, and so on up to A-th person.
// You need to tell the maximum and the minimum money that the ship company can earn if all A passengers buy tickets.

// Problem Constraints
// 1 ≤ A ≤ 3000
// 1 ≤ B ≤ 1000
// 1 ≤ C[i] ≤ 1000
// It is guaranteed that there are at least A empty seats in total.

// Input Format
// First argument is a integer A denoting the number of passengers in the queue.
// Second arugument is a integer B deonting the number of ships.
// Third argument is an integer array C of size B where C[i] denotes the number of empty seats in the i-th ship before the ticket office starts selling tickets.

// Output Format
// Return an array of size 2 denoting the maximum and minimum money that the ship company can earn.

// Example Input
// Input 1:

//  A = 4
//  B = 3
//  C = [2, 1, 1]
// Input 2:

//  A = 4
//  B = 3
//  C = [2, 2, 2]

// Example Output
// Output 1:

//  [5, 5]
// Output 2:

//  [7, 6]

// Example Explanation
// Explantion 1:

//  Maximum money can be earned if the passenger choose : 2(first ship) + 1(first ship) + 1(second ship) + 1(third ship).
//  So, the cost will be 5.
//  Minimum money can be earned if the passenger choose : 1(senocd ship) + 2(first ship) + 1(first ship) + 1(third ship).
//  So, the cost will be 5.
// Explanation 2:

//  Maximum money can be earned if the passenger choose : 2(first ship) + 2(second ship) + 2(third ship) + 1(first ship).
//  So, the cost will be 7.
//  Minimum money can be earned if the passenger choose : 2(senocd ship) + 2(first ship) + 1(first ship) + 1(second ship).
//  So, the cost will be 6.
func BuyTickets(A int, B int, C []int) []int {
	ans := make([]int, 0)

	person := A

	max := &maxHeap{}
	heap.Init(max)

	for i := 0; i < len(C); i++ {
		heap.Push(max, C[i])
	}

	maxAns := 0

	for person > 0 {
		x := heap.Pop(max).(int)
		maxAns += x
		x--
		if x > 0 {
			heap.Push(max, x)
		}
		person--
	}

	minAns := 0
	min := &minHeap1{}
	heap.Init(min)
	person = A
	for i := 0; i < len(C); i++ {
		heap.Push(min, C[i])
	}

	for person > 0 {
		x := heap.Pop(min).(int)
		minAns += x
		x--
		if x > 0 {
			heap.Push(min, x)
		}
		person--
	}

	ans = append(ans, maxAns)

	return ans
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type minHeap1 []int

func (h minHeap1) Len() int { return len(h) }
func (h minHeap1) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
