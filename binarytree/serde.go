package binarytree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	q := make([]*TreeNode, 0)

	q = append(q, root)

	ans := getString(root.Val) + ","

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				ans += getString(node.Left.Val) + ","
				q = append(q, node.Left)
			} else {
				ans += ","
			}

			if node.Right != nil {
				ans += getString(node.Right.Val) + ","
				q = append(q, node.Right)
			} else {
				ans += ","
			}
		}

	}

	ans = ans[:len(ans)-1]

	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	val, _ := getInt(arr[0])
	root := &TreeNode{Val: val}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	end := len(arr) - 1
	i := 1
	for i <= end && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		j := 0
		for j < 2 {
			if j == 0 {
				val, err := getInt(arr[i])
				if err != nil {
					node.Left = nil
				} else {
					node.Left = &TreeNode{Val: val}
					queue = append(queue, node.Left)
				}
				i++
				j++
			} else {
				val, err := getInt(arr[i])
				if err != nil {
					node.Right = nil
				} else {
					node.Right = &TreeNode{Val: val}
					queue = append(queue, node.Right)
				}
				i++
				j++
			}
		}

	}

	return root
}

func getInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	return val, err
}

func getString(i int) string {
	return strconv.Itoa(i)
}

// ================================================================================================================

// Serializes a tree to a single string.
func (c *Codec) serializeOpt(root *TreeNode) string {
	if root == nil {
		return ""
	}
	arr := []*TreeNode{root}
	for idx := 0; idx < len(arr); idx++ {
		v := arr[idx]
		if v == nil {
			continue
		}
		arr = append(arr, v.Left, v.Right)
	}

	out := make([]string, len(arr))
	for i := range arr {
		if arr[i] == nil {
			out[i] = "_"
		} else {
			out[i] = strconv.Itoa(arr[i].Val)
		}
	}
	return strings.Join(out, ".")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserializeOpt(data string) *TreeNode {
	if data == "" {
		return nil
	}

	arr := strings.Split(data, ".")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(arr[0])

	q := []*TreeNode{root}
	for slow, fast := 0, 1; fast < len(arr); slow, fast = slow+1, fast+2 {
		n := q[slow]
		if v, err := strconv.Atoi(arr[fast]); err == nil {
			n.Left = &TreeNode{Val: v}
			q = append(q, n.Left)
		}
		if v, err := strconv.Atoi(arr[fast+1]); err == nil {
			n.Right = &TreeNode{Val: v}
			q = append(q, n.Right)
		}
	}

	return root
}
