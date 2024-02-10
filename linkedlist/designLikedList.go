package linkedlist

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{Val: -1}
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.Next

	for i := 0; i < index; i++ {
		if curr != nil {
			curr = curr.Next
		} else {
			return -1
		}
	}

	if curr != nil {
		return curr.Val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{Val: val}
	newNode.Next = this.Next
	this.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	curr := this
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this
	for i := 0; i < index; i++ {
		if curr.Next != nil {
			curr = curr.Next
		} else {
			return
		}
	}

	newNode := &MyLinkedList{Val: val}
	newNode.Next = curr.Next
	curr.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	curr := this
	for i := 0; i < index; i++ {
		if curr.Next != nil {
			curr = curr.Next
		} else {
			return
		}
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
