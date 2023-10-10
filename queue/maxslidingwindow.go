package queue

func MaxSlidingWindow(A []int, k int) []int {

	// max := math.MinInt
	// idx := 0
	q := NewDeque(A[0])
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		if (i != 0 && i < len(A)) && A[i] < A[i-1] {
			q.PushBack(A[i])
		} else {
			for !q.IsEmpty() && q.Rear() < A[i] {
				q.RemoveBack()
			}
			if i != 0 {
				if !q.IsEmpty() {
					q.PushBack(A[i])
				} else {
					q = NewDeque(A[i])
				}
			}
		}
	}

	res = append(res, q.Top())
	if A[0] == q.Top() {
		q.RemoveFront()
	}

	// // q.Enqueue(idx)

	s := 1
	e := k

	for e < len(A) {
		if A[e] < A[e-1] {
			if q.IsEmpty() {
				q = NewDeque(A[e])
			} else {
				q.PushBack(A[e])
			}
		} else {
			i := e
			for !q.IsEmpty() && q.Rear() < A[i] {
				q.RemoveBack()
			}
			if q.IsEmpty() {
				q = NewDeque(A[e])
			} else {
				q.PushBack(A[e])
			}
		}

		res = append(res, q.Top())
		if A[s] == q.Top() {
			q.RemoveFront()
		}

		s++
		e++
	}

	return res
}
