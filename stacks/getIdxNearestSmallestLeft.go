package stacks

func GetNearestSmallestLeftIdx(A []int) []int {
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
