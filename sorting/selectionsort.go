package sorting

import "fmt"

func SelectionSort(arr []int) {
	arr = []int{8, 1, 3, 6, 11, 2, 4, 9, 7, 6} //1,2,3,4,6,6,7,8,9,11

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}

	fmt.Print(arr)
}
