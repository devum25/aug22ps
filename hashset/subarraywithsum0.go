package hashset

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func CheckForSubArrayWithSumZERO() bool { // with help of hashset
	A := []int{2, 2, 1, -3, 4, 3, 1, -2, -3}

	// pref [2,4,5,2,6,9,10,8,5]

	set := hashset.New()

	prefSum := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if i > 0 {
			prefSum[i] = A[i] + prefSum[i-1]
		} else {
			prefSum[0] = A[i]
		}
	}

	duplicate := hashset.New()

	for i := 0; i < len(prefSum); i++ {
		if !set.Contains(prefSum[i]) {
			set.Add(prefSum[i])
		} else {
			duplicate.Add(prefSum[i])
		}
	}

	if duplicate.Size() > 0 {
		return true
	}
	fmt.Print(prefSum)

	return false
}

func GetForSubArrayWithSumZERO() [][]int { // loop through prefix sum array, also returning indexes of subarrays
	A := []int{2, 2, 1, -3, 4, 3, 1, -2, -3}

	// pref [2,4,5,2,6,9,10,8,5]

	prefSum := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if i > 0 {
			prefSum[i] = A[i] + prefSum[i-1]
		} else {
			prefSum[0] = A[i]
		}
	}

	x := make(map[int]int)

	res := make([][]int, 0)

	for i := 0; i < len(prefSum); i++ {

		val, ok := x[prefSum[i]]
		if !ok {
			x[prefSum[i]] = i
		} else {
			v := []int{val + 1, i}
			res = append(res, v)
		}
	}

	return res
}

func GetIfSumKExistsForTwoIndex() bool { // O(N2)

	A := []int{2, 2, 1, -3, 4, 3, 1, -2, -3}

	k := 6

	for i := 0; i < len(A); i++ {

		for j := i + 1; j < len(A); j++ {

			if A[i]+A[j] == k {
				return true
			}
		}

	}

	return false
}

func GetIfSumKExistsForTwoIndexON() bool { // O(N)

	A := []int{2, 2, 1, -3, 4, 3, 1, -2, -3}

	k := 6

	set := hashset.New()

	for i := 0; i < len(A); i++ {
		if set.Contains(k - A[i]) {
			return true
		} else {
			set.Add(A[i])
		}
	}

	return false
}

func GetIfDifferenceKExistsForTwoIndexON() bool { // O(N) i<j A[i]-A[j] = k, A[i] = k+A[j]
	// for every i find k+A[i] value
	A := []int{2, 2, 1, -3, 4, 3, 1, 2, -3}
	k := 2
	// A := []int{12, 7, 11, 9, 7}

	// k := 3

	set := hashset.New()

	for i := 0; i < len(A); i++ {
		x := k + A[i]
		if set.Contains(x) {
			return true
		} else {
			set.Add(A[i])
		}
	}

	return false
}

// func CountNumberOfPairForSumK() int {
// 	A := []int{2, 2, 1, -3, 4, 3, 1, -2, -3}

// 	k := 4

// 	for i := 0; i < len(A); i++ {

// 	}

// 	return 0
// }
