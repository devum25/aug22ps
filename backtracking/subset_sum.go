package backtracking

import "fmt"

func CountSubsetWithSumK(arr []int, k int) int {
	lst := make([]int, 0)
	// ans := 0
	ans := make([][]int, 0)
	count := 0
	countSubset(arr, 0, &lst, k, &count, &ans)
	fmt.Println(ans)
	return count
}

func countSubset(arr []int, currIdx int, lst *[]int, k int, count *int, ans *[][]int) {
	if currIdx == len(arr) {
		sum := 0
		temp := make([]int, len(*lst))

		x := len(*lst)
		for i := 0; i < x; i++ {
			y := *lst
			temp[i] = y[i]
			sum += y[i]
		}

		if sum == k {
			x := *count
			x++
			*count = x
			*ans = append(*ans, temp)
		}
		return
	}

	*lst = append(*lst, arr[currIdx])
	countSubset(arr, currIdx+1, lst, k, count, ans)

	y := *lst
	y = y[:len(y)-1]
	*lst = y
	countSubset(arr, currIdx+1, lst, k, count, ans)
}

func countSubsetOptimal(arr []int, currIdx int, sum *int, k int, count *int) {
	if currIdx == len(arr) {
		if *sum == k {
			*count++
		}
		return

	}

	*sum += arr[currIdx]
	countSubsetOptimal(arr, currIdx+1, sum, k, count)

	*sum -= arr[currIdx]
	countSubsetOptimal(arr, currIdx+1, sum, k, count)
}
