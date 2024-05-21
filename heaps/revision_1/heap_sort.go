package revision1

func HeapSort(arr []int) {
	buildMaxHeap(arr, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func buildMaxHeap(arr []int, n int) {
	startIdx := (n / 2) - 1

	for i := startIdx; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i // initialize largest as root
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[largest] < arr[left] {
		largest = left
	}

	if right < n && arr[largest] < arr[right] {
		largest = right
	}

	if largest == i {
		return
	}

	arr[largest], arr[i] = arr[i], arr[largest]

	heapify(arr, n, largest)
}
