package queue

func GetNthNumber(n int) int {
	// c := 0
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	t := 0
	for n > 0 {
		n--
		t = q.Dequeue()
		q.Enqueue(t*10 + 1)
		q.Enqueue(t*10 + 2)
		q.Enqueue(t*10 + 3)

	}

	return t

}
