package subarrays

func AllSubArrays(arr []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		x := make([]int, 0)
		for j := i; j < len(arr); j++ {
			x = getSubArray(arr, i, j)
			res = append(res, x)
		}

	}
	return res
}

func getSubArray(arr []int, i, j int) []int {
	x := make([]int, 0)
	for k := i; k <= j; k++ {
		x = append(x, arr[k])
	}

	return x
}
