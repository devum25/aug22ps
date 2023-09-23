package stacks

// func SolveStak(A []int) []int {
// 	x := &MyStack{}
// 	stack := x.NewMyStack()
// 	min := math.MaxInt
// 	for i := 0; i < len(A); i++ {
// 		k := i
// 		min = math.MaxInt
// 		for j := i; j < len(A); j++ {
// 			if A[j] < min {
// 				min = A[j]
// 				k = j
// 			}
// 		}

// 		A[i], A[k] = A[k], A[i]
// 		stack.Push(A[i])

// 	}

// 	res := make([]int, len(A))

// 	for i := len(A) - 1; i >= 0; i-- {
// 		res[i] = stack.Pop()
// 	}

// 	return res
// }

// type MyStack struct {
// 	items []int
// }

// func (s *MyStack) NewMyStack() *MyStack {
// 	return &MyStack{items: []int{}}
// }

// func (s *MyStack) Push(x int) {
// 	s.items = append(s.items, x)
// }

// func (s *MyStack) Pop() int {
// 	val := s.items[len(s.items)-1]
// 	s.items = s.items[0 : len(s.items)-1]
// 	return val
// }

// func (s *MyStack) Peek() int {
// 	if s.Len() > 1 {
// 		val := s.items[len(s.items)-1]
// 		return val
// 	}
// 	return -1
// }

// func (s *MyStack) Len() int {
// 	return len(s.items)
// }
