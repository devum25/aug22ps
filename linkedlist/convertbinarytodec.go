package linkedlist

import "math"

func GetDecimalValue2(head *Node) (decimal int) {

	for head != nil {
		decimal = (decimal * 2) + head.Val
		head = head.Next
	}

	return
}

func GetDecimalValue(head *Node) int {
	h := head

	l := reverseLL(h)
	ans := 0
	c := 0
	for l != nil {
		ans = ans + int(math.Pow(float64(2), float64(c)))*l.Val
		c++
		l = l.Next
	}

	return ans
}

func reverseLL(head *Node) *Node {
	h := head
	var h1 *Node

	for h != nil {
		tmp := h
		h = h.Next
		tmp.Next = h1
		h1 = tmp
	}

	return h1
}
