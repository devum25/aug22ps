package darraygo

func Rotate90(arr [][]int) {
	reverse(arr)
}
func getTranspose(A [][]int) {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			A[i][j], A[j][i] = A[j][i], A[i][j]
		}
	}
}

func reverse(A [][]int) {

	for i := 0; i < len(A); i++ {
		k := len(A) - 1
		for j := 0; j < len(A); j++ {
			if j < k {
				A[i][j], A[i][k] = A[i][k], A[i][j]
				k--
			}
		}
	}
}
