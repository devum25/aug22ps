package stacks

import (
	"github.com/golang-collections/collections/stack"
)

func MinLength(s string) int {
	s = "ACBBD"
	x := "AB"
	y := "CD"

	var chars1 = []rune(x)
	var chars2 = []rune(y)

	stack := stack.New()
	// var last byte

	for i := 0; i < len(s); i++ {

		if stack.Len() == 0 {
			stack.Push(s[i])
			continue
		}

		if (stack.Peek().(byte) == byte(chars1[0]) && s[i] == byte(chars1[1])) || (stack.Peek().(byte) == byte(chars2[0]) && s[i] == byte(chars2[1])) {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}

	}

	var res []byte

	for i := stack.Len(); i > 0; i-- {
		res = append(res, stack.Pop().(byte))
	}

	return len(res)
}
