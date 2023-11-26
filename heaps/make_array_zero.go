package heaps

// 2357. Make Array Zero by Subtracting Equal Amounts
// Easy
// 1.1K
// 44
// Companies
// You are given a non-negative integer array nums. In one operation, you must:

// Choose a positive integer x such that x is less than or equal to the smallest non-zero element in nums.
// Subtract x from every positive element in nums.
// Return the minimum number of operations to make every element in nums equal to 0.

func MakeArrayZero(arr []int) int {
	heap := NewPairMinHeap()
	if arr[0] > 0 {
		heap.Insert(0, arr[0])
	}
	count := 0
	for i := 1; i < len(arr); i++ { // O(n)
		if len(heap.Items) == 0 && arr[i] > 0 {
			heap.Insert(i, arr[i])
		} else if len(heap.Items) > 0 && arr[i] < heap.Items[0][0] && arr[i] > 0 {
			heap.DeleteMin()
			heap.Insert(i, arr[i])
		}
	}

	for len(heap.Items) > 0 {
		temp := heap.DeleteMin()
		for i := 0; i < len(arr); i++ {
			if arr[i] > 0 {
				x := arr[i] - temp[0]
				arr[i] = x
				if len(heap.Items) == 0 && x > 0 {
					heap.Insert(i, x)
				} else if x > 0 && x < heap.Items[0][0] {
					heap.DeleteMin()
					heap.Insert(i, x)
				}
			}
		}
		count++
	}
	return count
}

// ======================================================================================

// Another optimal solution
// count number of distinct elements
func minimumOperations(nums []int) (c int) {
	// unique table
	v := make([]*struct{}, 101)
	v[0] = &struct{}{}

	// walkthrough each element of numbers
	for _, n := range nums {
		// if current number already set
		// skip it
		if v[n] != nil {
			continue
		}

		// count and set number
		v[n] = &struct{}{}
		c++
	}
	return
}
