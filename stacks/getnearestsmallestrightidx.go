package stacks

func GetNearestSmallestRightIdx(A []int) []int {
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
