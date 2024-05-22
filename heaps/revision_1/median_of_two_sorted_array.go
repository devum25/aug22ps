package revision1

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := merge(nums1, nums2)

	n := len(arr)

	if n%2 == 0 {
		x := n / 2
		ans := float64((float64(arr[x-1]) + float64(arr[x])) / float64(2))
		return ans
	}
	x := n / 2
	ans := float64(arr[x])
	return ans
}

func merge(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))

	i := 0
	j := 0
	c := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[c] = arr1[i]
			i++
			c++
		} else {
			arr[c] = arr2[j]
			j++
			c++
		}
	}

	if i < len(arr1) {
		for i < len(arr1) {
			arr[c] = arr1[i]
			i++
			c++
		}
	}

	if j < len(arr2) {
		for j < len(arr2) {
			arr[c] = arr2[j]
			j++
			c++
		}
	}

	return arr
}
