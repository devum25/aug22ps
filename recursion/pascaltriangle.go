package recursion

func PascalTriangle(n int) [][]int {
	ans := make([][]int, 0)
	helppascal(n-1, &ans)
	return ans
}

func helppascal(n int, ans *[][]int) {
	y := *ans
	if n == 0 {
		y = append(y, []int{1})
		*ans = y
		return
	}

	y = append(y, []int{1})
	*ans = y
	helppascal(n-1, ans)

	x := *ans
	y = *ans
	for i := 0; i < (len(x[n-1]) - 1); i++ {
		val := x[n-1][i] + x[n-1][i+1]
		y[n] = append(y[n], val)
	}

	y[n] = append(y[n], 1)
	*ans = y

}
