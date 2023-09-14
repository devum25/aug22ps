package stacks

import "github.com/golang-collections/collections/stack"

func RemoveConsecutiveDuplicate(s string) string {
	stack := stack.New()

	var last byte

	for i := len(s) - 1; i >= 0; i-- { //acb l  b c k
		if stack.Len() == 0 {
			stack.Push(s[i])
			last = s[i]
			continue
		}
		if s[i] != last {
			stack.Push(s[i])
			last = s[i]
		} else {
			stack.Pop()
			// stack.Push(s[i])
			if stack.Len() > 0 {
				x := stack.Peek()
				if x != nil {
					last = x.(byte)
				}
			}
		}

	}

	res := make([]byte, 0)

	for i := stack.Len(); i > 0; i-- {
		// if stack.Len() > 0 {
		res = append(res, stack.Pop().(byte))
		// }
	}

	return string(res)
}
