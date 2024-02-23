package arrayquestion

// Given an integer array nums,
// return an integer array counts where counts[i] is the number of smaller elements to the right of nums[i].
// https://www.youtube.com/watch?v=_sA1xI4XK0c
func CountSmaller(nums []int) []int {
	v := make([][2]int, len(nums))
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		v[i] = [2]int{nums[i], i}
	}

	smaller(v, count, 0, len(nums)-1)

	return count
}

func smaller(arr [][2]int, count []int, l, r int) [][2]int {
	if l == r {
		return arr
	}

	mid := (l + r) / 2
	smaller(arr, count, l, mid)
	smaller(arr, count, mid+1, r)

	return counting(arr, count, l, mid, r)
}

func counting(arr [][2]int, count []int, l, m, r int) [][2]int {
	temp := make([][2]int, r-l+1)

	i := l
	j := m + 1
	k := 0

	for i <= m && j <= r {
		if arr[i][0] > arr[j][0] {
			count[arr[i][1]] += (r - j + 1)
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= m {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= r {
		temp[k] = arr[j]
		j++
		k++
	}
	k = 0
	for i := l; i <= r; i++ {
		arr[i] = temp[k]
		k++
	}

	return arr
}
