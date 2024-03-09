package queue

type MyCircularQueue struct {
	Items []int
	f     int
	r     int
	size  int
}

func NewConstructor(k int) MyCircularQueue {
	return MyCircularQueue{Items: make([]int, k), f: -1, r: -1}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == len(this.Items) {
		return false
	}
	this.r = (this.r + 1) % len(this.Items)
	this.Items[this.r] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	this.f = (this.f + 1) % len(this.Items)
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.Items[(this.f+1)%len(this.Items)]
}

func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	return this.Items[(this.r)%len(this.Items)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.size == len(this.Items) {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
