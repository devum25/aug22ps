package heaps

func KPlaceApart(A []int, B int) []int {
	if len(A) == 0 {
		return nil
	} else if len(A) < B {
		return nil
	}

	res := make([]int, len(A))

	heap := NewMinHeap()

	k := B
	i := 0
	for k >= 0 {
		heap.Insert(A[i])
		i++
		k--
	}

	j := i
	n := 0
	for j < len(A) {
		res[n] = heap.DeleteMin()
		heap.Insert(A[j])
		j++
		n++
	}

	for len(heap.Items) > 0 {
		res[n] = heap.DeleteMin()
		n++
	}

	return res
}
