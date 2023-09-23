package stacks

func FinalPrices(prices []int) []int {
	stack := NewNumStack()
	res := make([]int, len(prices))

	stack.Push(prices[len(prices)-1])
	res[len(prices)-1] = stack.Top()

	for i := len(prices) - 2; i >= 0; i-- {

		for stack.Len() > 0 && stack.Top() >= prices[i] {
			stack.Pop()
		}

		if stack.Len() > 0 {
			res[i] = prices[i] - stack.Top()
			stack.Push(prices[i])
		} else {
			res[i] = prices[i]
			stack.Push(prices[i])
		}

	}

	return res
}

type NumStack struct {
	items []int
}

func NewNumStack() *NumStack {
	return &NumStack{items: []int{}}
}

func (ns *NumStack) Push(x int) {
	ns.items = append(ns.items, x)
}

func (ns *NumStack) Pop() int {
	val := ns.items[len(ns.items)-1]
	ns.items = ns.items[0 : len(ns.items)-1]
	return val
}

func (ns *NumStack) Top() int {
	return ns.items[len(ns.items)-1]
}

func (nc *NumStack) Len() int {
	return len(nc.items)
}
