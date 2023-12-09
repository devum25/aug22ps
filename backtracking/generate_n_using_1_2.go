package backtracking

import "fmt"

func Generate_1(n int) [][]int {
	ans := make([][]int, 0)
	lst := make([]int, n)
	generate(n, 0, &lst, &ans)
	fmt.Println(len(ans))
	return ans
}

func generate(n int, idx int, lst *[]int, ans *[][]int) {
	if idx == n {
		temp := make([]int, len(*lst))

		for i := 0; i < len(temp); i++ {
			h := *lst
			temp[i] = h[i]
		}
		*ans = append(*ans, temp)
		return
	}

	x := *lst
	x[idx] = 1
	*lst = x
	generate(n, idx+1, lst, ans)

	y := *lst
	y[idx] = 2
	*lst = y
	generate(n, idx+1, lst, ans)
}
