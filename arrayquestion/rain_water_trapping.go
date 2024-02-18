package arrayquestion

func Trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	left := 0
	for i := 0; i < len(height); i++ {
		leftMax[i] = left

		if height[i] > left {
			left = height[i]
		}
	}

	right := 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMax[i] = right

		if height[i] > right {
			right = height[i]
		}
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		x := min(leftMax[i], rightMax[i]) - height[i]
		if x > 0 {
			ans = ans + x
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
