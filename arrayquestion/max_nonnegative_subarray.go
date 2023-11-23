package arrayquestion

import "math"

// Given an array of integers, A of length N, find out the maximum sum sub-array of non negative numbers from A.
// The sub-array should be contiguous i.e., a sub-array created by choosing the second and fourth element and skipping the third element is invalid.
// Maximum sub-array is defined in terms of the sum of the elements in the sub-array.
// Find and return the required subarray.
// NOTE:

//     1. If there is a tie, then compare with segment's length and return segment which has maximum length.
//     2. If there is still a tie, then return the segment with minimum starting index.

// Input Format:
// The first and the only argument of input contains an integer array A, of length N.
// Output Format:
// Return an array of integers, that is a subarray of A that satisfies the given conditions.
// Constraints:
// 1 <= N <= 1e5
// -INT_MAX < A[i] <= INT_MAX
// Examples:
// Input 1:
//     A = [1, 2, 5, -7, 2, 3]
// Output 1:
//     [1, 2, 5]
// Explanation 1:
//     The two sub-arrays are [1, 2, 5] [2, 3].
//     The answer is [1, 2, 5] as its sum is larger than [2, 3].
// Input 2:
//     A = [10, -1, 2, 3, -4, 100]
// Output 2:
//     [100]
// Explanation 2:
//     The three sub-arrays are [10], [2, 3], [100].
//     The answer is [100] as its sum is larger than the other two.

func MaxSet(arr []int) []int {
	max := math.MinInt

	s := 0
	e := 0

	ans := make([]int, 3)
	curr := 0
	for e < len(arr) {
		if arr[e] >= 0 {
			curr = curr + arr[e]
			if curr > max {
				max = curr
				ans[0] = s
				ans[1] = e
				ans[2] = max
			} else if curr == max {
				temp := e - s + 1
				size := arr[1] - arr[0] + 1
				if temp > size {
					ans[0] = s
					ans[1] = e
				}
			}
		} else if arr[e] < 0 {
			s = e + 1
			curr = 0
		}

		e++
	}

	res := make([]int, 0)
	for i := ans[0]; i <= ans[1]; i++ {
		if arr[i] >= 0 {
			res = append(res, arr[i])
		}
	}

	return res
}
