package heaps

// Given an array of integers, A denoting a stream of integers. New arrays of integer B and C are formed.
// Each time an integer is encountered in a stream, append it at the end of B and append the median of array B at the C.
// Find and return the array C.
// NOTE:
// If the number of elements is N in B and N is odd, then consider the median as B[N/2] ( B must be in sorted order).
// If the number of elements is N in B and N is even, then consider the median as B[N/2-1]. ( B must be in sorted order).

func RunnningMedian(arr []int) []int {
	min := NewMinHeap()
	max := NewMaxHeap()
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(min.Items) == 0 && len(max.Items) == 0 {
			max.Insert(arr[i])
			res[0] = max.Items[0]
		} else {
			if arr[i] < max.Items[0] {
				max.Insert(arr[i])
				if len(max.Items)-len(min.Items) > 1 {
					min.Insert(max.DeleteMax())
				}
				if len(max.Items) == len(min.Items) {
					res[i] = max.Items[0]
				} else {
					if len(max.Items) > len(min.Items) {
						res[i] = max.Items[0]
					} else {
						res[i] = min.Items[0]
					}
				}
			} else {
				min.Insert(arr[i])
				if len(min.Items)-len(max.Items) > 1 {
					max.Insert(min.DeleteMin())
				}
				if len(max.Items) == len(min.Items) {
					res[i] = max.Items[0]
				} else {
					if len(max.Items) > len(min.Items) {
						res[i] = max.Items[0]
					} else {
						res[i] = min.Items[0]
					}
				}
			}
		}

	}

	return res
}
