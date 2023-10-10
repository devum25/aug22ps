package queue

type Deque struct {
	items *ListNode // this will be doubly linked list
}

func NewDeque(v int) *Deque {
	return &Deque{items: NewListNode(v)}
}

func (q *Deque) PushFront(x int) { // Insert front
	node := q.items.InsertAtStart(x)
	q.items = node
}

func (q *Deque) PushBack(x int) { // Enqueue
	q.items.InsertAtEnd(x)
}

func (q *Deque) RemoveFront() int { //Dequeue
	val := q.items.Top()
	q.items = q.items.RemoveFromStart()
	return val
}

func (q *Deque) RemoveBack() int { // remove from rear
	val := q.items.Rear()
	q.items = q.items.RemoveFromEnd()
	return val
}

func (q *Deque) Top() int {
	return q.items.val
}

func (q *Deque) Rear() int {
	return q.items.last.val
}

func (q *Deque) size() int {
	return q.items.Size()
}

func (q *Deque) IsEmpty() bool {
	return q.items.IsEmpty()
}
