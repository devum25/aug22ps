package arrayquestion

func IsIdealPermutation(nums []int) bool {
	a1 := make([]int, len(nums))
	a2 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a1[i] = nums[i]
		a2[i] = nums[i]
	}
	return inversionCount(a1) == localCount(a2)
}

func localCount(arr []int) int {
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			count++
		}
	}
	return count
}

func inversionCount(arr []int) int {
	var count int
	inversions1(arr, 0, len(arr)-1, &count)
	return count
}

func inversions1(arr []int, l, r int, count *int) []int {
	if l == r {
		return arr
	}

	mid := (l + r) / 2
	inversions1(arr, l, mid, count)
	inversions1(arr, mid+1, r, count)

	return countInversions1(arr, count, l, mid, r)
}

func countInversions1(arr []int, c *int, l, m, r int) []int {
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
			*c += (m - i + 1)

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
