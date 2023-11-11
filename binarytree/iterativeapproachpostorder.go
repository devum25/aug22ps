package binarytree

// Three types of Iterative Postorder Traversals.

// Using 1 Stack. O(n) Time & O(n) Space
// This is similar to Inorder using 1 Stack. The difference is we keep track of the previously printed node in pre. And we only print a node if its right child is null or equal to pre.
// Push all left nodes into the stack till it hits NULL.
// root = s.peek()
// if root.right = null or pre (Means we have traversed the right subtree already)
// We print root and pop it from s.
// Make pre = root
// root = null (So we dont go down to left child again)
// else
// root = root.right (Traverse the right subtree before printing root)
// Keep iterating till both the below conditions are met -
// Stack is empty and
// Root is NULL.

func IterativePostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := NewStack()

	curr := root
	pre := root
	res := make([]int, 0)

	for curr != nil || !stack.IsEmpty() {
		if curr != nil {
			stack.Push(curr)
			curr = curr.Left
		} else {
			temp := stack.Top()
			if temp.Right != nil && temp.Right != pre {
				curr = temp.Right
			} else {
				node := stack.Pop()
				res = append(res, node.Val)
				pre = node
				curr = nil
			}
		}
	}

	return res

}
