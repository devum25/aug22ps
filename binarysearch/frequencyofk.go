package binarysearch

func FrequencyOfElement(A []int, B int) int {

	first := firstOccurence(A, B)
	second := secondOccurence(A, B)

	if first >= 0 && second >= 0 {
		return second - first + 1
	}

	return 0
}

func firstOccurence(A []int, k int) int {
	l := 0
	r := len(A) - 1
	idx := -1
	for l <= r {
		mid := (l + r) / 2

		if A[mid] == k {
			idx = mid
			r = mid - 1
		} else if A[mid] < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return idx
}

func secondOccurence(A []int, k int) int {
	l := 0
	r := len(A) - 1
	idx := -1
	for l <= r {
		mid := (l + r) / 2

		if A[mid] == k {
			idx = mid
			l = mid + 1
		} else if A[mid] < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return idx
}
