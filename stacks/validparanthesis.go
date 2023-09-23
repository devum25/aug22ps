package stacks

/**
 * @input A : String
 *
 * @Output Integer
 */
func solve(s string) int {
	ms := &MyStack1{}
	st := ms.NewMyStack()
	if len(s) == 1 || len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if st.Len() == 0 {
			st.Push(s[i])
			continue
		}
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			st.Push(s[i])

			continue
		} else {
			if st.Len() > 0 {
				if st.Peek() == '(' && s[i] == ')' {
					st.Pop()
				} else if st.Peek() == '{' && s[i] == '}' {
					st.Pop()
				} else if st.Peek() == '[' && s[i] == ']' {
					st.Pop()
				} else {
					return 0
				}
			}
		}
	}
	if st.Len() == 0 {
		return 1
	}

	return 0
}

type MyStack1 struct {
	items []byte
}

func (s *MyStack1) NewMyStack() *MyStack1 {
	s.items = make([]byte, 0)
	return s
}

func (s *MyStack1) Push(x byte) {
	s.items = append(s.items, x)
}

func (s *MyStack1) Pop() byte {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *MyStack1) Peek() byte {
	val := s.items[len(s.items)-1]
	return val
}

func (s *MyStack1) Len() int {
	return len(s.items)
}
