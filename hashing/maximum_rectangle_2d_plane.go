package hashset

import "strconv"

func MaximumRectangle(A [][2]int) int {
	hash := make(map[string]bool)

	for i := 0; i < len(A); i++ {
		s := getString(A[i][0]) + ":" + getString(A[i][1])
		hash[s] = true
	}
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i][0] != A[j][0] && A[i][1] != A[j][1] {
				s1 := getString(A[i][0]) + ":" + getString(A[j][1])
				s2 := getString(A[j][0]) + ":" + getString(A[i][1])

				if hash[s1] && hash[s2] {
					count++
				}
			}
		}
	}

	return count / 4
}

func getString(a int) string {
	return strconv.Itoa(a)
}
