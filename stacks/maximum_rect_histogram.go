package stacks

func LargestRectangleArea1(A []int) int {
	nsl := getNearestSmallestLeft(A)
	nsr := getNearestSmallestRight(A)

	ans := 0
	for i := 0; i < len(A); i++ {
		if nsl[i] == -1 && nsr[i] == -1 {
			ans = max(ans, A[i]*len(A))
		} else if nsl[i] == -1 {
			ans = max(ans, A[i]*((i+1)+(nsr[i]-i-1)))
		} else if nsr[i] == -1 {
			ans = max(ans, A[i]*((len(A)-i)+(i-nsl[i]-1)))
		} else {
			w := nsr[i] - nsl[i] - 1
			ans = max(ans, w*A[i])
		}
	}
	return ans
}

func getNearestSmallestLeft(A []int) []int {
	res := make([]int, len(A))

	stack := NewStack1()

	res[0] = -1
	stack.Push(0)

	for i := 1; i < len(A); i++ {

		for stack.Len() > 0 && A[stack.Peek()] >= A[i] {
			stack.Pop()
		}
		if stack.Len() == 0 {
			res[i] = -1
		} else {
			res[i] = stack.Peek()
		}
		stack.Push(i)
	}

	return res
}

func getNearestSmallestRight(A []int) []int {
	res := make([]int, len(A))

	stack := NewStack1()

	res[len(A)-1] = -1
	stack.Push(len(A) - 1)

	for i := len(A) - 2; i >= 0; i-- {

		for stack.Len() > 0 && A[stack.Peek()] >= A[i] {
			stack.Pop()
		}
		if stack.Len() == 0 {
			res[i] = -1
		} else {
			res[i] = stack.Peek()
		}
		stack.Push(i)
	}

	return res
}

type Stack5 struct {
	Items []int
}

func NewStack5() *Stack5 {
	return &Stack5{Items: make([]int, 0)}
}

func (s *Stack5) Push(x int) {
	s.Items = append(s.Items, x)
}

func (s *Stack5) Pop() int {
	var x int
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
	}
	return x
}

func (s *Stack5) Peek() int {
	var x int
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
	}
	return x
}

func (s *Stack5) Len() int {
	return len(s.Items)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
