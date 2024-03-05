package stacks

//https://www.youtube.com/watch?v=j313ttNJjo0
func RemoveDuplicateLetters1(s string) string {
	stack := NewStack()

	freq := make([]int, 26)
	visited := make([]bool, 26)

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a'] = freq[s[i]-'a'] + 1
	}

	stack.Push(s[0])
	visited[s[0]-'a'] = true
	freq[s[0]-'a'] = freq[s[0]-'a'] - 1

	for i := 1; i < len(s); i++ {
		if !visited[s[i]-'a'] {
			if int(s[i]-'a') > int(stack.Peek()-'a') {
				stack.Push(s[i])
				visited[s[i]-'a'] = true
				freq[s[i]-'a'] = freq[s[i]-'a'] - 1
			} else {
				for stack.Len() > 0 && int(stack.Peek()-'a') > int(s[i]-'a') && freq[stack.Peek()-'a'] > 0 {
					visited[stack.Peek()-'a'] = false
					visited[s[i]-'a'] = true
					stack.Pop()
				}
				stack.Push(s[i])
			}
		}
	}
	var res string
	for stack.Len() > 0 {
		res += string(stack.Pop())
	}

	return reverse(res)
}

func reverse(s string) string {
	var x string
	for i := len(s) - 1; i >= 0; i-- {
		x += string(s[i])
	}

	return x
}

type MyStack2 struct {
	Items []byte
}

func NewStack() *MyStack2 {
	return &MyStack2{Items: make([]byte, 0)}
}

func (s *MyStack2) Push(x byte) {
	s.Items = append(s.Items, x)
}

func (s *MyStack2) Pop() byte {
	var x byte
	if len(s.Items) > 0 {
		x = s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
	}

	return x
}

func (s *MyStack2) Peek() byte {
	if len(s.Items) > 0 {
		return s.Items[len(s.Items)-1]
	}
	return '0'
}

func (s *MyStack2) Len() int {
	return len(s.Items)
}
