package stacks

import (
	"math"

	"github.com/golang-collections/collections/stack"
)

// Implement getMin() of stack function extending library
type MyStack struct {
	MinStack *stack.Stack
	*stack.Stack
	min int
}

func MyNewStack() *MyStack {
	return &MyStack{Stack: stack.New(), MinStack: stack.New(), min: math.MaxInt}
}

func (ms *MyStack) GetMin() int {
	return ms.min
}

func (ms *MyStack) Push(x int) {
	if x <= ms.min {
		ms.min = x
		ms.MinStack.Push(x)
	}

	ms.Stack.Push(x)

}

func (ms *MyStack) Pop() int {
	val := ms.Stack.Pop().(int)

	if val == ms.min {
		ms.MinStack.Pop()
		ms.min = ms.MinStack.Peek().(int)
	}

	return val
}
