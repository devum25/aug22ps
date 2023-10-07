package queue

type MyQueue struct {
	stack1 MyStack
	stack2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{stack1: *NewStack(), stack2: *NewStack()}
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	for !this.stack1.IsEmpty() {
		this.stack2.Push(this.stack1.Pop())
	}

	val := this.stack2.Pop()

	for !this.stack2.IsEmpty() {
		this.stack1.Push(this.stack2.Pop())
	}

	return val
}

func (this *MyQueue) Peek() int {
	for !this.stack1.IsEmpty() {
		this.stack2.Push(this.stack1.Pop())
	}

	val := this.stack2.Peek()

	for !this.stack2.IsEmpty() {
		this.stack1.Push(this.stack2.Pop())
	}

	return val
}

func (this *MyQueue) Empty() bool {
	return this.stack1.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyStack struct {
	items []int
}

func NewStack() *MyStack {
	return &MyStack{items: []int{}}
}

func (s *MyStack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *MyStack) Pop() int {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *MyStack) Peek() int {
	val := s.items[len(s.items)-1]
	return val
}

func (s *MyStack) IsEmpty() bool {
	return len(s.items) == 0
}
