package queue

type Queue struct {
	items *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(x int) {
	if q.items == nil {
		q.items = newLinkedList(x)
	} else {
		q.items.insertAtEnd(x)
	}
}

func (q *Queue) Dequeue() int {
	var val int
	if q.items != nil {
		val, q.items = q.items.removeFromHead()
		return val
	}
	return 0
}

type Node struct {
	val  int
	next *Node
}

func newLinkedList(v int) *Node {
	return &Node{val: v}
}

func (n *Node) insertAtEnd(x int) *Node {
	h := n
	for h.next != nil {
		h = h.next
	}
	h.next = &Node{val: x}
	return n
}

func (n *Node) removeFromHead() (int, *Node) {
	if n != nil {
		val := n.val
		n = n.next
		return val, n
	}

	return 0, nil
}
