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

func (q *Queue) Top() int {
	if q.items != nil {
		return q.items.val
	}
	return -1
}

func (q *Queue) size() int {
	return q.items.size()
}

type Node struct {
	val    int
	next   *Node
	last   *Node
	length int
}

func newLinkedList(v int) *Node {
	return &Node{val: v, length: 1}
}

func (n *Node) insertAtEnd(x int) *Node {
	h := n
	if n.last == nil {
		h.next = &Node{val: x}
		n.last = h.next
	} else {
		n.last.next = &Node{val: x}
		n.last = n.last.next
	}
	n.length++

	return n
}

func (n *Node) removeFromHead() (int, *Node) {
	if n != nil {
		last := n.last
		l := n.length
		val := n.val
		n = n.next
		if n != nil {
			n.last = last
			n.length = l - 1
		}
		return val, n
	}

	return 0, nil
}

func (n *Node) size() int {
	return n.length
}
