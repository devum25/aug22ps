package binarysearch

func FindPeak(A []int) int { // element which is greater than it's neighbour, local minima
	if len(A) == 1 {
		return 0
	} else if len(A) == 0 {
		return -1
	}

	if A[0] > A[1] {
		return 0
	} else if A[len(A)-1] > A[len(A)-2] {
		return len(A) - 1
	}

	l := 1
	r := len(A) - 2

	for l <= r {
		mid := (l + r) / 2

		if A[mid-1] < A[mid] && A[mid+1] < A[mid] {
			return mid
		} else if A[mid-1] > A[mid] {
			r = mid - 1
		} else if A[mid+1] > A[mid] {
			l = mid + 1
		}
	}

	return -1
}
