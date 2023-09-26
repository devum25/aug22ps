package stacks

func NextGreater(A []int) []int {
	stack := NewIntStack()
	stack.Push(A[len(A)-1])
	res := make([]int, len(A))
	res[len(A)-1] = -1
	for i := len(A) - 2; i >= 0; i-- {

		for stack.Len() > 0 && stack.Peek() <= A[i] {
			stack.Pop()
		}
		if stack.Len() > 0 {
			res[i] = stack.Peek()
			stack.Push(A[i])
		} else {
			res[i] = -1
			stack.Push(A[i])
		}
	}

	return res
}

type IntStack struct {
	items []int
}

func NewIntStack() *IntStack {
	return &IntStack{items: []int{}}
}

func (ns *IntStack) Push(x int) {
	ns.items = append(ns.items, x)
}

func (ns *IntStack) Pop() int {
	val := ns.items[len(ns.items)-1]
	ns.items = ns.items[0 : len(ns.items)-1]
	return val
}

func (ns *IntStack) Peek() int {
	return ns.items[len(ns.items)-1]
}

func (nc *IntStack) Len() int {
	return len(nc.items)
}
