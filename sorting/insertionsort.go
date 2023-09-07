package sorting

import "fmt"

func InsertionSort(A []int) {
	A = []int{8, 1, 3, 6, 11, 2, 4, 9, 7, 6} //1,2,3,4,6,6,7,8,9,11

	for i := 1; i < len(A); i++ {
		j := i
		k := i
		for j > 0 {
			if A[j-1] > A[k] {
				A[k], A[j-1] = A[j-1], A[k]
				k--
			}
			j--
		}
	}

	fmt.Print(A)
}
