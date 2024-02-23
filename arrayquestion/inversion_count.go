package arrayquestion

func InversionCount(arr []int) int {
	var count int
	inversions(arr, 0, len(arr)-1, &count)
	return count
}

func inversions(arr []int, l, r int, count *int) []int {
	if l == r {
		return arr
	}

	mid := (l + r) / 2
	inversions(arr, l, mid, count)
	inversions(arr, mid+1, r, count)

	return countInversions(arr, count, l, mid, r)
}

func countInversions(arr []int, c *int, l, m, r int) []int {
	temp := make([]int, r-l+1)

	i := l
	j := m + 1
	k := 0

	for i <= m && j <= r {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
			k++
		} else {
			temp[k] = arr[j]
			j++
			k++
			if m == i {
				*c += 1
			} else {
				*c += (m - i)
			}
		}
	}

	if i <= m {
		for i <= m {
			temp[k] = arr[i]
			i++
			k++
		}
	}

	if j <= r {
		for j <= r {
			temp[k] = arr[j]
			j++
			k++
		}
	}
	k = 0
	for x := l; x <= r; x++ {
		arr[x] = temp[k]
		k++
	}

	return arr
}
