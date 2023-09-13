package linkedlist

import "math"

func LongestPalindrome(head *Node) int {
	curr := head
	var prev *Node
	ans := 0

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		ans = int(math.Max(float64(ans), float64(2*count(prev, next)+1)))
		ans = int(math.Max(float64(ans), float64(2*count(curr, next))))
		prev = curr
		curr = next
	}

	return ans
}

func count(A *Node, B *Node) int {
	x := A
	y := B
	sum := 0
	for x != nil && y != nil {
		if x.Val != y.Val {
			break
		}
		sum++
		x = x.Next
		y = y.Next
	}

	return sum
}
