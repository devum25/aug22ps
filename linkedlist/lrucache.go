package linkedlist

type ListNode struct {
	key  int
	val  int
	prev *ListNode
	next *ListNode
}

type LRUCache struct {
	m    map[int]*ListNode
	cap  int
	len  int
	head *ListNode
	tail *ListNode
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*ListNode, capacity),
		cap:  capacity,
		len:  0,
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}

	if node == this.tail {
		return node.val
	}

	this.adjustElements(node)

	return node.val
}

func (this *LRUCache) adjustElements(node *ListNode) {
	h := node
	for h != nil && h.next != nil {
		next := h.next
		prev := h.prev
		tmp := next.next
		if prev != nil {
			prev.next = next
		}
		next.prev = prev
		next.next = h
		h.next = tmp
		h.prev = next
	}

	this.tail = h

	x := h
	c := h
	for x != nil {
		c = x
		x = x.prev
	}

	this.head = c
}

func (this *LRUCache) Put(key int, val int) {
	n := &ListNode{key: key, val: val, next: nil, prev: nil}
	if node, ok := this.m[key]; ok {
		node.val = val
	}
	if this.len == 1 && this.cap == 1 {
		delete(this.m, this.head.key)
		this.head = n
		this.tail = n
		this.m[n.key] = n
		return
	}

	if this.len >= this.cap {
		delete(this.m, this.head.key)
		h := this.head.next
		h.prev = nil
		this.head = h
		this.tail.next = n
		n.prev = this.tail
		this.tail = n
		this.m[n.key] = n

	} else {
		if this.tail != nil {
			this.tail.next = n
			n.prev = this.tail
			this.tail = n
			this.m[n.key] = n
			this.len++
		} else {
			this.head = n
			this.tail = n
			this.m[n.key] = n
			this.len++
		}
	}
}
