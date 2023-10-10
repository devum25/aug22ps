package queue

type MyStackNew struct {
	q1 QueueNew
	q2 QueueNew
}

func ConstructorNew() MyStackNew {
	return MyStackNew{q1: *NewQueueNew(), q2: *NewQueueNew()}
}

func (this *MyStackNew) Push(x int) {
	this.q1.Enqueue(x)
}

func (this *MyStackNew) Pop() int {
	for this.q1.items != nil && this.q1.size() > 1 {
		this.q2.Enqueue(this.q1.Dequeue())
	}

	val := this.q1.Dequeue()

	for this.q2.items != nil {
		this.q1.Enqueue(this.q2.Dequeue())
	}

	return val
}

func (this *MyStackNew) Top() int {
	for this.q1.items != nil && this.q1.size() > 1 {
		this.q2.Enqueue(this.q1.Dequeue())
	}

	val := this.q1.Dequeue()

	for this.q2.items != nil {
		this.q1.Enqueue(this.q2.Dequeue())
	}

	this.q1.Enqueue(val)

	return val
}

func (this *MyStackNew) Empty() bool {
	return this.q1.items == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type QueueNew struct {
	items *Node
}

func NewQueueNew() *QueueNew {
	return &QueueNew{}
}

func (q *QueueNew) Enqueue(x int) {
	if q.items == nil {
		q.items = newLinkedList(x)
	} else {
		q.items.insertAtEnd(x)
	}
}

func (q *QueueNew) Dequeue() int {
	var val int
	if q.items != nil {
		val, q.items = q.items.removeFromHead()
		return val
	}
	return 0
}

func (q *QueueNew) Top() int {
	if q.items != nil {
		return q.items.val
	}
	return -1
}

func (q *QueueNew) size() int {
	return q.items.size()
}
