package revision1

import "math"

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

/////////////////////////////////////////////////////////////////////////////////////////////////////////

// O(log(N+M)) APPROACH

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1) // always run on smaller array
	}
	low := 0
	high := n1
	left := (n1 + n2 + 1) / 2
	n := n1 + n2

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := left - mid1
		l1 := math.MinInt
		l2 := math.MinInt
		r1 := math.MaxInt
		r2 := math.MaxInt
		if mid1 < n1 {
			r1 = nums1[mid1]
		}
		if mid2 < n2 {
			r2 = nums2[mid2]
		}
		if (mid1 - 1) >= 0 {
			l1 = nums1[mid1-1]
		}
		if (mid2 - 1) >= 0 {
			l2 = nums2[mid2-1]
		}
		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return max(l1, l2)
			}
			return float64((max(l1, l2) + min(r1, r2)) / 2.0)
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

func min(l1, l2 int) float64 {
	if l1 > l2 {
		return float64(l2)
	}
	return float64(l1)
}

func max(l1, l2 int) float64 {
	if l1 > l2 {
		return float64(l1)
	}
	return float64(l2)
}
