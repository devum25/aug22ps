package queue

type ListNode struct {
	val    int
	next   *ListNode
	prev   *ListNode
	last   *ListNode
	length int
}

func NewListNode(v int) *ListNode {
	l := &ListNode{val: v}
	l.last = l
	l.length++
	return l
}

func (l *ListNode) InsertAtEnd(x int) *ListNode {
	if l == nil {
		l = &ListNode{val: x}
		l.last = l
		l.length++
		return l
	}
	l.last.next = &ListNode{val: x}
	l.last.next.prev = l.last
	l.last = l.last.next
	l.length++
	return l
}

func (l *ListNode) InsertAtStart(x int) *ListNode {
	h := &ListNode{val: x}
	l.prev = h
	h.next = l
	h.last = l.last
	l.last = nil
	h.length = l.length + 1
	return h
}

func (l *ListNode) RemoveFromEnd() *ListNode {
	prev := l.last.prev
	if prev != nil {
		prev.next = nil
		l.last = prev
	} else {
		l = nil
		return l
	}
	l.length--
	return l
}

func (l *ListNode) RemoveFromStart() *ListNode {
	h := l.next
	if h != nil {
		h.prev = nil
		if l.last == l {
			l.last = nil
			h.last = h
		} else {
			h.last = l.last
		}
		l.next = nil
		h.length = l.length - 1
	}
	return h
}

func (l *ListNode) Top() int {
	if l != nil {
		return l.val
	}

	return -1
}

func (l *ListNode) Rear() int {
	if l != nil {
		return l.last.val
	}

	return -1
}

func (l *ListNode) IsEmpty() bool {
	return l == nil
}

func (l *ListNode) Size() int {
	return l.length
}
