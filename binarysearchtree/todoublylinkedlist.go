package binarysearchtree

type ListNode struct {
	Left  *ListNode
	Value int
	Right *ListNode
}

func treeNode_new(val int) *ListNode {
	var node *ListNode = new(ListNode)
	node.Value = val
	node.Left = nil
	node.Right = nil
	return node
}

func ConvertBSTtoDLL(node *ListNode) *ListNode {

	converter(node)
	return head
}

var prev *ListNode
var head *ListNode

func converter(root *ListNode) {
	if root == nil {
		return
	}

	converter(root.Left)
	if prev == nil {
		head = root
	} else {
		root.Left = prev
		prev.Right = root
	}

	prev = root

	converter(root.Right)

}
