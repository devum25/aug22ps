package stacks

func RemoveDuplicateLetters(s string) string {
	lastIdxs := getLastIndexArray(s)
	seen := make([]int, 26)
	stack := NewMyStack()
	stack.Push(s[0])
	seen[s[0]-'a'] = 1
	for i := 1; i < len(s); i++ {
		for stack.Len() > 0 && s[i] < stack.Peek() && lastIdxs[stack.Peek()-'a'] > i && seen[s[i]-'a'] == 0 {
			x := stack.Pop()
			seen[x-'a'] = 0
		}
		if seen[s[i]-'a'] == 0 {
			stack.Push(s[i])
			seen[s[i]-'a'] = 1
		}

	}

	res := make([]byte, stack.Len())

	for i := stack.Len(); i > 0; i-- {
		res[i-1] = stack.Pop()
	}

	return string(res)
}

func getLastIndexArray(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for i := 0; i < len(s); i++ {
		res[s[i]-'a'] = i
	}
	return res
}

type MyStack struct {
	items []byte
}

func NewMyStack() *MyStack {
	return &MyStack{items: make([]byte, 0)}
}

func (s *MyStack) Push(x byte) {
	s.items = append(s.items, x)
}

func (s *MyStack) Pop() byte {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *MyStack) Peek() byte {
	val := s.items[len(s.items)-1]
	return val
}

func (s *MyStack) Len() int {
	return len(s.items)
}
