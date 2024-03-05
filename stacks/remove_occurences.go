package stacks

func RemoveOccurrences(s string, part string) string {
	stack := NewStack()
	var str string
	idx := 0
	if s[0] == part[0] {
		str += string(s[0])
		idx++
	}
	stack.Push(s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == part[idx] {
			str += string(s[i])
			idx++
			stack.Push(s[i])
			if len(str) == len(part) { // pop out from stack
				c := 0
				for c < len(part) {
					stack.Pop()
					c++
					idx--
				}
				if stack.Len() > 0 && stack.Peek() == part[idx] {
					idx++
				}
			}
		} else {
			stack.Push(s[i])
			idx = 0
			str = ""
		}
	}

	var res string

	for stack.Len() > 0 {
		res += string(stack.Pop())
	}

	return reverse(res)
}

func reverse1(s string) string {
	var x string

	for i := len(s) - 1; i >= 0; i-- {
		x += string(s[i])
	}

	return x
}

type MyStack4 struct {
	Items []byte
}

func NewStack4() *MyStack4 {
	return &MyStack4{Items: make([]byte, 0)}
}

func (s *MyStack4) Push(x byte) {
	s.Items = append(s.Items, x)
}

func (s *MyStack4) Pop() byte {
	var x byte
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
	}

	return x
}

func (s *MyStack4) Peek() byte {
	var x byte
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
	}
	return x
}

func (s *MyStack4) Len() int {
	return len(s.Items)
}
