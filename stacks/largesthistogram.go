package stacks

func LargestRectangleArea(heights []int) int {
	if len(heights) == 1 {
		return heights[0]
	}
	left := nearestSmallestLeftIdx(heights)
	right := newSmallestRightIdx(heights)
	max := 0
	for i := 0; i < len(heights); i++ { // 2 4 6 9
		if right[i] == -1 {
			x := len(heights) - 1
			right[i] = x - i + 1
		}

		if left[i] == -1 && i != 0 {

			left[i] = i

		}
		dist := 0
		if left[i] == -1 && i == 0 {
			dist = right[i]
		} else if left[i] == 0 {
			dist = right[i]
		} else {
			dist = right[i] - left[i] - 1
		}

		val := dist * heights[i]
		if val > max {
			max = val
		}
	}

	return max
}

func getDistance(a, b, n int) int {
	return 0
}

func nearestSmallestLeftIdx(A []int) []int {

	stack := NewNumStack()

	res := make([]int, len(A))
	n := len(A) - 1
	res[0] = -1
	stack.Push(0)

	for i := 1; i <= n; i++ {

		for stack.Len() > 0 && A[i] <= A[stack.Top()] {
			stack.Pop()
		}

		if stack.Len() > 0 {
			res[i] = stack.Top()
			stack.Push(i)
		} else {
			res[i] = -1
			stack.Push(i)
		}

	}

	return res
}

func newSmallestRightIdx(A []int) []int {
	stack := NewNumStack()

	res := make([]int, len(A))
	n := len(A) - 1
	res[n] = -1
	stack.Push(n)

	for i := n - 1; i >= 0; i-- {

		for stack.Len() > 0 && A[i] <= A[stack.Top()] {
			stack.Pop()
		}

		if stack.Len() > 0 {
			res[i] = stack.Top()
			stack.Push(i)
		} else {
			res[i] = -1
			stack.Push(i)
		}

	}

	return res
}
