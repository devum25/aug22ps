package sorting

func QuickSort(A []int) []int {
	quick(A, 0, len(A)-1)
	return A
}

func quick(A []int, s, e int) {
	if s >= e {
		return
	}
	pivot := (s + e) / 2
	A[pivot], A[s] = A[s], A[pivot]
	idx := partition(A, s, e)
	quick(A, s, idx-1)
	quick(A, idx+1, e)
}

func partition(A []int, l, r int) int {
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

	return e
}
