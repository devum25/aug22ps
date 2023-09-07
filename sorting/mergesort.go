package sorting

import "fmt"

func Mergesort(arr []int) {
	mergeSorting(arr, 0, len(arr)-1)
	fmt.Print(arr)
}

func mergeSorting(A []int, l, r int) {

	// assumption
	//mergeSorting func will sort the subarray from l to r
	if l == r {
		return
	}

	mid := (l + r) / 2
	mergeSorting(A, l, mid)
	mergeSorting(A, mid+1, r)
	mergeSortedArray(A, l, mid+1, r)

}

func mergeSortedArray(arr []int, l, m, r int) {
	arr1 := make([]int, m-l)
	arr2 := make([]int, r-m+1)

	j := 0
	for i := l; i < m; i++ {
		arr1[j] = arr[i]
		j++
	}

	j = 0
	for i := m; i <= r; i++ {
		arr2[j] = arr[i]
		j++
	}

	i := 0
	j = 0

	x := l
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[x] = arr1[i]
			i++
			x++
		} else {
			arr[x] = arr2[j]
			j++
			x++
		}
	}

	if i < len(arr1) {

		for i < len(arr1) {
			arr[x] = arr1[i]
			x++
			i++
		}

	}
	if j < len(arr2) {
		for i < len(arr2) {
			arr[x] = arr2[i]
			x++
			i++
		}
	}

	fmt.Print(arr)
}
