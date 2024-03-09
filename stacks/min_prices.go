package stacks

func FinalPrices1(prices []int) []int {
	stack := NewStack1()
	minStack := NewStack1()
	res := make([]int, len(prices))

	for i := len(prices) - 1; i >= 0; i-- {
		if stack.Len() == 0 {
			stack.Push(prices[i])
			minStack.Push(prices[i])
			res[i] = prices[i]
		} else {
			if stack.Peek() <= prices[i] {
				res[i] = prices[i] - stack.Peek()
				if prices[i] <= stack.Peek() {
					minStack.Push(prices[i])
				}
				stack.Push(prices[i])
			} else if minStack.Peek() <= prices[i] {
				res[i] = prices[i] - minStack.Peek()
				if prices[i] <= stack.Peek() {
					minStack.Push(prices[i])
				}
				stack.Push(prices[i])
			} else {
				res[i] = prices[i]
				stack.Push(prices[i])
				minStack.Push(prices[i])
			}
		}
	}

	return res
}

type Stack1 struct {
	Items []int
}

func NewStack1() *Stack1 {
	return &Stack1{Items: make([]int, 0)}
}

func (s *Stack1) Push(x int) {
	s.Items = append(s.Items, x)
}

func (s *Stack1) Pop() int {
	var x int
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
	}

	return x
}

func (s *Stack1) Peek() int {
	var x int
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
	}

	return x
}

func (s *Stack1) Len() int {
	return len(s.Items)
}
