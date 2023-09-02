package hashset

import "fmt"

func SubArraySum() int {
	// A := []int{7, 9, 4, 5, 9}
	// k := 18

	A := []int{1, 2, 3} //[1,3,6]
	k := 3
	result := 0

	pfs := make([]int, len(A))

	for s := 0; s < len(A); s++ {
		if s != 0 {
			pfs[s] = pfs[s-1] + A[s]
		} else {
			pfs[s] = A[s]
		}
	}

	for s := 0; s < len(pfs); s++ {
		e := s
		for ; e < len(pfs); e++ {
			if s == 0 {
				sum := pfs[e]
				if sum == k {
					result++
				}
			} else {
				sum := pfs[e] - pfs[s-1]
				if sum == k {
					result++
				}
			}
		}
	}
	fmt.Print(result)
	return result
}
