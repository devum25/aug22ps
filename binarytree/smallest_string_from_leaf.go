package binarytree

import "sort"

func SmallestFromLeaf(root *TreeNode) string {
	lst := make([]int, 0)
	curr := 0
	p := 1
	solve5(root, &lst, &curr, &p)

	sort.Ints(lst)

	return getString1(lst[0])
}

func solve5(root *TreeNode, lst *[]int, curr, p *int) {
	if root == nil {
		*lst = append(*lst, *curr)
		return
	}

	*curr = (*p)*root.Val + (*curr)
	*p = (*p) * 10

	//if root.Left != nil {
	solve5(root.Left, lst, curr, p)
	//}
	*p = (*p) / 10
	*curr = (*curr) - (*p)*root.Val

	//if root.Right != nil {
	solve5(root.Right, lst, curr, p)
	//}
}

func getString1(num int) string {
	str := make([]byte, 0)
	for num > 0 {
		temp := num / 10
		str = append(str, byte(temp+int('a')))
		num = num % 10
	}

	return reverse1(string(str))
}

func reverse1(str string) string {
	s := make([]byte, 0)

	for i := len(str) - 1; i >= 0; i-- {
		s = append(s, str[i])
	}

	return string(s)
}
