package arrayquestion

func SortArray(nums []int) []int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(arr []int, l, r int) []int {
	if l == r {
		return arr
	}

	mid := (l + r) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	return merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) []int {
	a := make([]int, r-l+1)
	c := 0
	i := 0
	j := m + 1
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			a[c] = arr[i]
			i++
			c++
		} else {
			a[c] = arr[j]
			j++
			c++
		}
	}

	if i <= m {
		for i <= m {
			a[c] = arr[i]
			i++
			c++
		}
	}

	if j <= r {
		for j <= r {
			a[c] = arr[j]
			j++
			c++
		}
	}

	return a
}
