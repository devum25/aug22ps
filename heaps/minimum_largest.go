package heaps

import "math"

// Problem Description
// Given an array A of N numbers, you have to perform B operations. In each operation, you have to pick any one of the N elements and add the original value(value stored at the index before we did any operations) to its current value. You can choose any of the N elements in each operation.

// Perform B operations in such a way that the largest element of the modified array(after B operations) is minimized.
// Find the minimum possible largest element after B operations.

// Problem Constraints
// 1 <= N <= 106
// 0 <= B <= 105
// -104 <= A[i] <= 104

// Input Format
// The first argument is an integer array A.
// The second argument is an integer B.

// Output Format
// Return an integer denoting the minimum possible largest element after B operations.

// Example Input
// Input 1:

//  A = [1, 2, 3, 4]
//  B = 3
// Input 2:

//  A = [5, 1, 4, 2]
//  B = 5

// Example Output
// Output 1:

//  4
// Output 2:

//  5

// Example Explanation
// Explanation 1:

//  Apply operation on element at index 0, the array would change to [2, 2, 3, 4]
//  Apply operation on element at index 0, the array would change to [3, 2, 3, 4]
//  Apply operation on element at index 0, the array would change to [4, 2, 3, 4]
//  Minimum possible largest element after 3 operations is 4.
// Explanation 2:

//  Apply operation on element at index 1, the array would change to [5, 2, 4, 2]
//  Apply operation on element at index 1, the array would change to [5, 3, 4, 2]
//  Apply operation on element at index 1, the array would change to [5, 4, 4, 2]
//  Apply operation on element at index 1, the array would change to [5, 5, 4, 2]
//  Apply operation on element at index 3, the array would change to [5, 5, 4, 4]
//  Minimum possible largest element after 5 operations is 5.

// maintain current state in an array
// next state of every element with index in min heap

// in every operation fetch from minheap, take original value based on index
// modify current state array and insert back to heap the modified index for next state

func MinLargest(arr []int, b int) int {
	min := NewPairMinHeap()

	curr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		min.Insert(i, arr[i]+arr[i])
		curr[i] = arr[i]
	}

	k := 1

	for k <= b {
		temp := min.DeleteMin()
		curr[temp[1]] = curr[temp[1]] + arr[temp[1]]
		min.Insert(temp[1], curr[temp[1]]+arr[temp[1]])
		k++
	}

	max := math.MinInt

	for i := 0; i < len(curr); i++ {
		if curr[i] > max {
			max = curr[i]
		}
	}

	return max
}

type PairMinHeap struct {
	Items [][2]int
	Size  int
}

func NewPairMinHeap() PairMinHeap {
	return PairMinHeap{Items: make([][2]int, 0)}
}

func (heap *PairMinHeap) Insert(idx, val int) {
	if len(heap.Items) == 0 {
		x := [2]int{val, idx}
		heap.Items = append(heap.Items, x)
	} else {
		x := [2]int{val, idx}
		heap.Items = append(heap.Items, x)
		heap.Size++
	}

	i := heap.Size

	for i > 0 {
		p := (i - 1) / 2

		if heap.Items[p][0] > heap.Items[i][0] {
			heap.Items[p], heap.Items[i] = heap.Items[i], heap.Items[p]
		} else {
			break
		}

		i = p
	}
}

func (heap *PairMinHeap) DeleteMin() [2]int {
	val := heap.Items[0]
	heap.Items[0], heap.Items[len(heap.Items)-1] = heap.Items[len(heap.Items)-1], heap.Items[0]
	heap.Items = heap.Items[:heap.Size]
	if heap.Size != 0 {
		heap.Size--
	}

	i := 0

	for i < len(heap.Items) {
		min := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(heap.Items) && heap.Items[l][0] < heap.Items[min][0] {
			min = l
		}

		if r < len(heap.Items) && heap.Items[r][0] < heap.Items[min][0] {
			min = r
		}

		if min == i {
			break
		}

		heap.Items[min], heap.Items[i] = heap.Items[i], heap.Items[min]
		i = min
	}

	return val
}
