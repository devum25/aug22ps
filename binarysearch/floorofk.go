package binarysearch

func GetFloorOfElement(A []int, k int) int {
	l, r := 0, len(A)-1

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == k {
			return A[mid]
		} else if l == r {
			if l == r {
				if A[r] < k {
					return A[r]
				} else {
					return A[r-1]
				}
			}
		} else if A[mid] < k {
			l = mid + 1
		} else if A[mid] > k {
			r = mid - 1
		}

	}

	if l > r {
		return A[r]
	}

	return -1
}
