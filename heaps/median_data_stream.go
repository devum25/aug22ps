package heaps

func Median(arr []int) []int {
	heap := NewMinHeap()

	meadian := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		heap.Insert(arr[i])
		temp := heap.Items
		HeapSort(temp)

		if i == 0 {
			meadian[i] = arr[i]
			continue
		}

		if len(temp)%2 == 0 {
			mid := (0 + (len(temp) - 1)) / 2
			meadian[i] = (temp[mid] + temp[mid+1]) / 2
		} else {
			mid := (0 + (len(temp) - 1)) / 2
			meadian[i] = temp[mid]
		}
	}

	return meadian
}

// ============================================================================================================================
type MedianFinder struct {
	min    MinHeap
	max    MaxHeap
	median float64
}

func Constructor() MedianFinder {
	return MedianFinder{min: NewMinHeap(), max: NewMaxHeap()}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.max.Items) == 0 && len(this.min.Items) == 0 {
		this.max.Insert(num)
		this.median = float64(num)
		return
	}

	if float64(num) < this.median {
		this.max.Insert(num)
		if (len(this.max.Items) - len(this.min.Items)) > 1 {
			this.min.Insert(this.max.DeleteMax())
		}
		if len(this.max.Items) > len(this.min.Items) {
			this.median = float64(this.max.Items[0])
		} else {
			this.median = (float64(this.max.Items[0]+this.min.Items[0]) / 2)
		}
	} else {
		this.min.Insert(num)
		if (len(this.min.Items) - len(this.max.Items)) > 1 {
			this.max.Insert(this.min.DeleteMin())
		}
		if len(this.min.Items) > len(this.max.Items) {
			this.median = float64(this.min.Items[0])
		} else {
			this.median = (float64(this.min.Items[0]+this.max.Items[0]) / 2)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	return this.median
}
