package binarysearchtree

import "fmt"

func IterativeInorder(root *TreeNode) {
	stack := NewStack()

	curr := root

	for curr != nil || !stack.IsEmpty() {

		if curr != nil {
			stack.Push(curr)
			curr = curr.Left
		} else {
			temp := stack.Pop()
			fmt.Print(temp.Val)
			curr = temp.Right
		}
	}
}

type Stack struct {
	items []*TreeNode
}

func NewStack() *Stack {
	return &Stack{items: []*TreeNode{}}
}

func (s *Stack) Push(x *TreeNode) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() *TreeNode {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) > 0
}
