package backtracking

import "fmt"

func Generate_2(N int) {
	count := 0
	// ans := make([][]int, 0)
	lst := make([]int, N)
	// generated_3(N, 0, &lst, &ans, &count)
	generate_2(N, &lst, 0, &count)
	fmt.Println(count)
	fmt.Println(len(ans))
}

func generate_2(N int, cl *[]int, idx int, count *int) {

	if idx == N {
		*count++
		fmt.Println(cl)
		return
	}

	for i := 1; i <= 5; i++ {
		(*cl)[idx] = i
		generate_2(N, cl, idx+1, count)
	}

}

func generated_3(n int, idx int, lst *[]int, ans *[][]int, count *int) {
	if idx == n {
		temp := make([]int, len(*lst))

		for i := 0; i < len(temp); i++ {
			h := *lst
			temp[i] = h[i]
		}
		*count++
		*ans = append(*ans, temp)
		return
	}

	x := *lst
	x[idx] = 1
	*lst = x
	generated_3(n, idx+1, lst, ans, count)

	y := *lst
	y[idx] = 2
	*lst = y
	generated_3(n, idx+1, lst, ans, count)

	g := *lst
	g[idx] = 3
	*lst = g
	generated_3(n, idx+1, lst, ans, count)

	h := *lst
	h[idx] = 4
	*lst = h
	generated_3(n, idx+1, lst, ans, count)

	i := *lst
	i[idx] = 5
	*lst = i
	generated_3(n, idx+1, lst, ans, count)
}
