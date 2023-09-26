package stacks

// [1,2,3,4,3,2]
func NextGreaterElements(nums []int) []int {
	x := make([]int, len(nums))
	stack := NewIntStack()
	for i := len(nums) - 1; i >= 0; i-- {
		pushed := false
		for stack.Len() > 0 && stack.Peek() <= nums[i] {
			stack.Pop()
		}
		if stack.Len() > 0 {
			x[i] = stack.Peek()
			stack.Push(nums[i])
		} else {
			for j := 0; j < i; j++ {
				if nums[j] > nums[i] {
					x[i] = nums[j]
					pushed = true
					stack.Push(nums[i])
					// stack.Push(x[i])
					break
				}
			}
			if !pushed {
				x[i] = -1
				stack.Push(nums[i])
			}
		}
	}

	return x
}
