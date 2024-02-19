package darraygo

func SearchKInSorted(A [][]int, B int) int {
	i := 0
	j := len(A[i]) - 1
	for j >= 0 && i < len(A) {
		if A[i][j] == B {
			x := ((i + 1) * 1009) + j + 1
			if j >= 1 {
				for j >= 1 && A[i][j-1] == B {
					j--
					x = ((i + 1) * 1009) + j + 1

				}
			}
			return x
		}
		if B > A[i][j] {
			i++
		} else {
			j--
		}
	}

	return -1
}
