package stacks

import (
	"math"
)

type MinStack struct {
	items    []int
	min      int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{items: []int{}, minStack: []int{}, min: math.MaxInt32}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	if val <= this.min {
		this.min = val
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	item := this.items[len(this.items)-1]
	if item == this.min {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
		if len(this.minStack) > 0 {
			this.min = this.minStack[len(this.minStack)-1]
		} else {
			this.min = math.MaxInt
			//this.minStack = append(this.minStack,item)
		}
	}
	this.items = this.items[0 : len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//==============================================================================================================================
