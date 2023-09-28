package binarysearch

func FindMinima(A []int) int {
	if len(A) == 1 {
		return A[0]
	} else if len(A) == 0 {
		return -1
	}

	if A[0] < A[1] {
		return A[0]
	} else if A[len(A)-1] < A[len(A)-1] {
		return A[len(A)-1]
	}

	l := 1
	r := len(A) - 2

	for l <= r {
		mid := (l + r) / 2

		if A[mid-1] > A[mid] && A[mid+1] > A[mid] {
			return A[mid]
		} else if A[mid-1] < A[mid] {
			r = mid - 1
		} else if A[mid+1] < A[mid] {
			l = mid + 1
		}
	}

	return -1
}
