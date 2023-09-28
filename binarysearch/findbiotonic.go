package binarysearch

func FindInBiotonic(A []int, B int) int {
	peak := FindPeak(A)
	idx := -1
	idx = searchLeftElement(A, 0, peak, B)
	if idx >= 0 {
		return idx
	} else {
		idx = searchRightElement(A, peak, len(A)-1, B)
		if idx >= 0 {
			return idx
		}
	}
	return -1
}

func searchRightElement(A []int, l, r, b int) int {

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == b {
			return mid
		} else if A[mid] < b {
			r = mid - 1
		} else if A[mid] > b {
			l = mid + 1
		}
	}

	return -1

}

func searchLeftElement(A []int, l, r, b int) int {

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == b {
			return mid
		} else if A[mid] > b {
			r = mid - 1
		} else if A[mid] < b {
			l = mid + 1
		}
	}

	return -1

}
