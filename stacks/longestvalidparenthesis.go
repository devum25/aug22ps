package stacks

// (()
// )()())
// )()())
// ()(()

func longestValidParentheses(s string) int {
	stack := NewMyStack()
	if s == "" {
		return 0
	}
	ans := 0
	res := 0
	for i := len(s) - 1; i >= 0; i-- {

		for stack.Len() > 0 && s[i] == ')' && stack.Peek() == '(' {
			stack.Pop()
			res = res + 2
			if stack.Len() == 0 {
				ans = res + ans
			}
		}

		if s[i] == '(' {
			stack.Push(s[i])
		}
	}

	return res
}
