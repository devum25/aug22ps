package hashset

// Given N points in 2D plane, count number of right angle triangles
// O(N2) approach
func CountOfRightTriangle(A [][2]int) int {
	count := 0
	hash := make(map[string]bool)

	for i := 0; i < len(A); i++ {
		s := getString(A[i][0]) + ":" + getString(A[i][1])
		hash[s] = true
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i][0] != A[j][0] && A[i][1] != A[j][1] {
				s1 := getString(A[i][0]) + ":" + getString(A[j][1]) // x1y2
				s2 := getString(A[j][0]) + ":" + getString(A[i][1]) // x2y1

				if hash[s1] {
					count++
				}
				if hash[s2] {
					count++
				}
			}
		}
	}

	return count
}

// O(N) approach

func CountOfRightTriangleOptimal(A [][2]int) int {
	count := 0
	hashX := make(map[int]int)
	hashY := make(map[int]int)
	for i := 0; i < len(A); i++ {
		hashX[A[i][0]]++
		hashY[A[i][1]]++
	}

	for i := 0; i < len(A); i++ {
		a := hashX[A[i][0]]
		b := hashY[A[i][1]]
		count += (a - 1) * (b - 1)
	}

	return count
}
