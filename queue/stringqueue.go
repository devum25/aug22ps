package queue

type StringQueue struct {
	items *StringNode
}

func NewStringQueue() *StringQueue {
	return &StringQueue{}
}

func (q *StringQueue) Enqueue(x string) {
	if q.items == nil {
		q.items = newStringLinkedList(x)
	} else {
		q.items.insertAtEnd(x)
	}
}

func (q *StringQueue) Dequeue() string {
	var val string
	if q.items != nil {
		val, q.items = q.items.removeFromHead()
		return val
	}
	return ""
}

func (q *StringQueue) Top() string {
	if q.items != nil {
		return q.items.val
	}
	return ""
}

type StringNode struct {
	val  string
	next *StringNode
	last *StringNode
}

func newStringLinkedList(v string) *StringNode {
	return &StringNode{val: v}
}

func (n *StringNode) insertAtEnd(x string) *StringNode {
	h := n
	if n.last == nil {
		h.next = &StringNode{val: x}
		n.last = h.next
	} else {
		n.last.next = &StringNode{val: x}
		n.last = n.last.next
	}

	return n
}

func (n *StringNode) removeFromHead() (string, *StringNode) {
	if n != nil {
		last := n.last
		val := n.val
		n = n.next
		if n != nil {
			n.last = last
		}

		return val, n
	}

	return "", nil
}
