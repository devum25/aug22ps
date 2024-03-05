package stacks

type ItemStack struct {
	items []byte
}

func (s *ItemStack) New() *ItemStack {
	s.items = []byte{}
	return s
}

func (s *ItemStack) Push(t byte) {
	s.items = append(s.items, t)
}

func (s *ItemStack) Pop() byte {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *ItemStack) Peek() byte {
	return s.items[len(s.items)-1]
}
