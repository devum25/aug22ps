package darraygo

func Transpose(arr [][]int) [][]int {
	if arr == nil {
		return nil
	}

	row := len(arr[0])
	col := len(arr)

	res := make([][]int, row)

	for j := 0; j < row; j++ {
		x := make([]int, 0)
		for i := 0; i < col; i++ {
			x = append(x, arr[i][j])
		}

		res[j] = x
	}

	return res
}
