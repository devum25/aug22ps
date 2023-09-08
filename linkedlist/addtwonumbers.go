package linkedlist

import (
	"math"
)

func AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	h1 := l1
	pow := 0
	sum1 := 0
	for h1 != nil {
		sum1 += (int(math.Pow10(pow)) * h1.Val)
		pow++
		h1 = h1.Next
	}

	h2 := l2
	pow = 0
	sum2 := 0

	for h2 != nil {
		sum2 += (int(math.Pow10(pow)) * h2.Val)
		pow++
		h2 = h2.Next
	}

	sum := sum1 + sum2
	// rev := 0
	// for sum != 0 {
	// 	rev = (rev * 10) + (sum % 10)
	// 	sum = sum / 10
	// }
	var newNode *Node
	var head *Node
	for sum != 0 {
		if newNode == nil {
			newNode = &Node{Val: sum % 10}
			head = newNode
		} else {
			newNode.Next = &Node{Val: sum % 10}
			newNode = newNode.Next
		}

		sum = sum / 10
	}

	return head
}
