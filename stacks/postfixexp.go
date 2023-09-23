package stacks

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToPostFix(s string) string {
	res := make([]byte, 0)

	st := &Stack{}
	stack := st.NewStack()

	oper := make(map[byte]int)
	oper['/'] = 4
	oper['*'] = 3
	oper['+'] = 2
	oper['-'] = 2
	oper['('] = -1
	oper[')'] = -1

	bracketFound := false //3+10*(3-4/2)+3
	prev := s[0]
	if oper[prev] != 0 {
		stack.Push(prev)
	} else {
		res = append(res, prev)
	}
	for i := 1; i < len(s); i++ {
		if bracketFound && oper[s[i]] > 0 {
			// stack.Push operators
			stack.Push(s[i])
			prev = s[i]
			continue
		}
		if _, ok := oper[s[i]]; !ok {
			if oper[s[i]] == 0 {
				if oper[prev] == 0 {
					res = append(res, s[i])
				} else {
					res = append(res, byte(' '))
					res = append(res, s[i])
				}
			}
		} else {
			if stack.IsEmpty() {
				stack.Push(s[i])
				prev = s[i]
				continue
			}
			if s[i] == '(' {
				stack.Push(s[i])
				bracketFound = true
				prev = s[i]
				continue
			} else if s[i] == ')' {
				bracketFound = false
				for j := stack.Len(); j > 0; j-- {
					if stack.Top() == '(' {
						stack.Pop()
						break
					}
					res = append(res, byte(' '))
					res = append(res, stack.Pop())
				}
				prev = s[i]
				continue
			}

			if !stack.IsEmpty() && oper[s[i]] <= oper[stack.Top()] {

				for k := stack.Len(); k > 0; k-- {
					if oper[stack.Top()] < oper[s[i]] {
						//res = append(res, byte(' '))
						//res = append(res, stack.Pop())
						stack.Push(s[i])
						break
					}
					res = append(res, byte(' '))
					res = append(res, stack.Pop())
				}
				if stack.Len() == 0 {
					stack.Push(s[i])
				}
			} else {
				stack.Push(s[i])
			}

		}
		prev = s[i]
	}

	for i := stack.Len(); i > 0; i-- {
		res = append(res, byte(' '))
		res = append(res, stack.Pop())
	}

	calculatePostFix(string(res))

	return string(res)
}

func calculatePostFix(s string) int {
	fmt.Print(s)
	// s = strings.TrimLeft(s)
	st := &Stack{}
	stacks := st.NewStringStack()
	num := 0
	oper := make(map[string]int)
	oper["/"] = 4
	oper["*"] = 3
	oper["+"] = 2
	oper["-"] = 2

	str := strings.Split(s, " ")

	for i := 0; i < len(str); i++ {
		if oper[str[i]] == 0 && str[i] != "" {
			stacks.Push(str[i])
		} else if str[i] != "" {
			k := 0
			s := []string{}
			for j := stacks.Len(); j > 0; j-- {
				if k < 2 {
					s = append(s, stacks.Pop())
					k++
				} else {
					break
				}
			}

			num = calculate(s, str[i])
			stacks.Push(strconv.Itoa(num))
		}
	}

	// if stacks.Len() > 0 {
	// 	for i := stacks.Len(); i > 0; i-- {
	v, _ := strconv.Atoi(stacks.Pop())
	num = v
	// 	}
	// }
	return num
}

func calculate(s []string, y string) int {
	var a, b int
	if len(s) > 0 {
		a, _ = strconv.Atoi(s[1])
	}
	if len(s) > 1 {
		b, _ = strconv.Atoi(s[0])
	}

	switch y {

	case "+":
		{
			return a + b
		}
	case "-":
		{
			return a - b
		}
	case "/":
		{
			return a / b
		}
	case "*":
		{
			return a * b
		}

	}
	return 0
}

type Stack struct {
	items []byte
}

func (s *Stack) NewStack() *Stack {
	return &Stack{items: []byte{}}
}

func (s *Stack) Push(x byte) {
	s.items = append(s.items, x)
}
func (s *Stack) Pop() byte {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}
func (s *Stack) Top() byte {
	return s.items[len(s.items)-1]
}
func (s *Stack) Len() int {
	return len(s.items)
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

type StringStack struct {
	items []string
}

func (s *Stack) NewStringStack() *StringStack {
	return &StringStack{items: []string{}}
}

func (s *StringStack) Push(x string) {
	s.items = append(s.items, x)
}
func (s *StringStack) Pop() string {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}
func (s *StringStack) Top() string {
	return s.items[len(s.items)-1]
}
func (s *StringStack) Len() int {
	return len(s.items)
}
func (s *StringStack) IsEmpty() bool {
	return len(s.items) == 0
}
