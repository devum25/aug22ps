package arrayquestion

// Given an array of N elements, Partition the array from l to r such that.
// all elements <= A[l] go to left of it
// all elements > A[l] go to right of it

func PartitionOnFirstElementFromLtoR(A []int, l, r int) []int {
	s := l + 1
	e := r

	for s <= e {
		if A[s] <= A[l] {
			s++
		} else if A[e] > A[l] {
			e--
		} else {
			A[s], A[e] = A[e], A[s]
		}
	}

	A[l], A[e] = A[e], A[l]

	return A
}
