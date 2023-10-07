package queue

func ReversK(A []int, B int) []int {
	q1 := NewQueue()
	for i := 0; i < len(A); i++ {
		q1.Enqueue(A[i])
	}

	res := make([]int, len(A))
	for i := 1; i <= B; i++ {
		res[B-i] = q1.Dequeue()
	}

	for i := B; i < len(A); i++ {
		res[i] = q1.Dequeue()
	}

	return res
}
