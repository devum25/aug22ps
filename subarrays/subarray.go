package subarrays

import "fmt"

// You are given an integer array A.

// Decide whether it is possible to divide the array into one or more subarrays of even length such that the first and last element of all subarrays will be even.

// Return "YES" if it is possible; otherwise, return "NO" (without quotes).

func EvenSubarrays() bool {
	A := []int{978, 847, 95, 729, 778, 586, 188, 782, 813, 870, 871, 940, 312, 693, 580, 101, 760, 837, 564, 633, 680, 155, 241, 374, 682, 290, 850, 601, 433, 922, 773, 959, 530, 290, 990, 50, 516, 409, 868, 131, 664, 851, 721, 880, 20, 450, 745, 387, 787, 823, 392, 242, 674, 347, 65, 135, 819, 324, 651, 678, 139, 940}

	if len(A)%2 == 0 && A[0]%2 == 0 && A[len(A)-1]%2 == 0 {
		return true
	}

	return false
}

// Given an integer array A containing N distinct integers, you have to find all the leaders in array A.
// An element is a leader if it is strictly greater than all the elements to its right side.
func LeaderElement() {
	A := []int{16, 17, 4, 3, 5, 2} // ans := [17,5,2]

	res := make([]int, 0)
	last := 0
	for j := len(A) - 1; j >= 0; j-- {
		if j == len(A)-1 {
			res = append(res, A[j])
			last = A[j]
			continue
		}

		if A[j] > last {
			last = A[j]
			res = append(res, last)
		}

	}

	fmt.Print(res)
}

// You have given a string A having Uppercase English letters.
// You have to find how many times subsequence "AG" is there in the given string.
// NOTE: Return the answer modulo 109 + 7 as the answer can be very large.

func SpecialSubsequence() {
	A := "ABCGAG" // seq = "AG", ans = 3

	count := 0
	ans := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 'G' {
			count++
		} else if A[i] == 'A' && count > 0 {
			ans = ans + count
		}
	}

	fmt.Print(ans)
}
