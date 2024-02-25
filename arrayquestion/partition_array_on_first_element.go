package arrayquestion

// Given an array of N elements, Partition the array such that.
// all elements <= A[0] go to left of it
// all elements > A[0] go to right of it

func PartitionOnFirstElement(A []int) []int {
	s := 1
	e := len(A) - 1

	for s <= e {
		if A[s] <= A[0] {
			s++
		} else if A[e] > A[0] {
			e--
		} else {
			A[s], A[e] = A[e], A[s]
		}
	}

	A[0], A[e] = A[e], A[0]

	return A
}
