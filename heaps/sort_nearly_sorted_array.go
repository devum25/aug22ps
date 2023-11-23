package heaps

// Give an array(nearly sorted array), Sort the array
// NEARLY SORTED : K-sorted (Every element is atmost k position away from it's sorted position)

func SortNearlySorted(arr []int, k int) []int {
	miheap := NewMinHeap()

	x := toMinHeap(arr[:k])
	miheap.Items = x
	miheap.size = len(x) - 1
	res := make([]int, len(arr))
	j := 0
	for i := k; i < len(arr); i++ {
		res[j] = miheap.DeleteMin()
		miheap.Insert(arr[i])
		j++
	}

	if j < len(arr) {
		for j < len(arr) {
			res[j] = miheap.DeleteMin()
			j++
		}
	}

	return res
}
