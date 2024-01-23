package binarysearch

func SearchInRotatedSortedSingleBinarySearch(A []int, k int) int {

	l := 0
	r := len(A) - 1

	for l <= r {
		mid := (l + r) / 2
		if A[mid] == k {
			return mid
		}
		if k >= A[0] {
			if A[mid] >= A[0] {
				if A[mid] < k {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else if A[mid] < A[0] {
				r = mid - 1
			}
		} else {
			if A[mid] < A[0] {
				if A[mid] < k {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

// ##########################################################################################################
// ##########################################################################################################
func SearchInRotatedSorted(A []int, k int) int {
	if len(A) == 1 {
		if A[0] == k {
			return 0
		}
		return -1
	}
	if A[0] < A[len(A)-1] {
		return searchfromltor(A, 0, len(A)-1, k)
	}

	minima := findLocalMinima(A)

	if k >= A[0] {
		return searchfromltor(A, 0, minima-1, k)
	} else {
		return searchfromltor(A, minima, len(A)-1, k)
	}

}

func findLocalMinima(A []int) int {
	pivot := -1
	l := 0
	r := len(A) - 1

	for l <= r {
		mid := (l + r) / 2

		if A[mid] < A[0] {
			pivot = A[mid]
			r = mid - 1
		} else if A[mid] >= A[0] {
			l = mid + 1
		}
	}

	return pivot
}

func searchfromltor(A []int, l, r, k int) int {

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == k {
			return mid
		} else if A[mid] > k {
			r = mid - 1
		} else if A[mid] < k {
			l = mid + 1
		}
	}

	return -1

}

func Search(A []int, k int) int {
	l := 0
	r := len(A) - 1

	for l < r {
		mid := (l + r) / 2
		if A[mid] == k {
			return mid
		}
		if A[mid] < A[0] {
			if k >= A[0] {
				r = mid - 1
			} else if k < A[0] {
				if A[mid] > k {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		} else if A[mid] >= A[0] {
			if k >= A[0] {
				if A[mid] < k {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
