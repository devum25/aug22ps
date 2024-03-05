package stacks

// Create stack from stack with Push, Pop and GetMin

type StackWithMin struct {
	Stack    *MyStackLib
	minStack *MyStackLib
}

func NewStackWithMin() *StackWithMin {
	return &StackWithMin{Stack: NewMyStackLib(), minStack: NewMyStackLib()}
}

func (s *StackWithMin) Push(x int) {
	if len(s.Stack.Items) == 0 {
		s.Stack.Items = append(s.Stack.Items, x)
		s.minStack.Items = append(s.minStack.Items, x)
	} else {
		if x < s.minStack.Peek() {
			s.minStack.Items = append(s.minStack.Items, x)
		}
		s.Stack.Items = append(s.Stack.Items, x)
	}
}

func (s *StackWithMin) Pop() int {
	x := -1
	if len(s.Stack.Items) > 0 {
		x = s.Pop()
	}
	if x == s.minStack.Peek() {
		s.minStack.Pop()
	}
	return x
}

func (s *StackWithMin) GetMin(x int) int {
	return s.minStack.Peek()
}

type MyStackLib struct {
	Items []int
}

func NewMyStackLib() *MyStackLib {
	return &MyStackLib{Items: make([]int, 0)}
}

func (s *MyStackLib) Push(x int) {
	s.Items = append(s.Items, x)
}

func (s *MyStackLib) Peek() int {
	return s.Items[len(s.Items)-1]
}

func (s *MyStackLib) Pop() int {
	x := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return x
}
