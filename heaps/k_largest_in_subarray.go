package heaps

// Problem Description
// Given an integer array B of size N.
// You need to find the Ath largest element in the subarray [1 to i], where i varies from 1 to N. In other words, find the Ath largest element in the sub-arrays [1 : 1], [1 : 2], [1 : 3], ...., [1 : N].
// NOTE: If any subarray [1 : i] has less than A elements, then the output should be -1 at the ith index.
// Problem Constraints
// 1 <= N <= 100000
// 1 <= A <= N
// 1 <= B[i] <= INT_MAX

// Input Format
// The first argument is an integer A.
// The second argument is an integer array B of size N.

// Output Format
// Return an integer array C, where C[i] (1 <= i <= N) will have the Ath largest element in the subarray [1 : i].

// Example Input
// Input 1:
//  A = 4
//  B = [1 2 3 4 5 6]
// Input 2:
//  A = 2
//  B = [15, 20, 99, 1]

// Example Output
// Output 1:
//  [-1, -1, -1, 1, 2, 3]
// Output 2:
//  [-1, 15, 20, 20]

func AthLargestInSubarray(A int, B []int) []int {
	min := NewMinHeap()
	res := make([]int, len(B))
	k := A
	i := 0
	for k > 1 {
		min.Insert(B[i])
		res[i] = -1
		i++
		k--
	}

	min.Insert(B[i])
	res[i] = min.Items[0]
	i++

	for j := i; j < len(B); j++ {
		if B[j] > min.Items[0] {
			min.DeleteMin()
			min.Insert(B[j])
			res[j] = min.Items[0]
		} else {
			res[j] = min.Items[0]
		}
	}

	return res
}
