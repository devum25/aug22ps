package binarysearch

func MaximumCount(A []int) int {
	if A[0] > 0 {
		return len(A)
	}
	if A[0] > 0 {
		return len(A)
	}
	if A[0] == 0 && A[len(A)-1] == 0 {
		return 0
	}
	if len(A) == 1 {
		if A[0] != 0 {
			return 1
		} else {
			return 0
		}
	}

	l := 0
	r := len(A) - 1
	n := -1
	p := -1
	for l <= r {
		mid := (l + r) / 2

		if A[mid] < 0 {
			if (mid+1) < len(A) && A[mid+1] >= 0 {
				n = mid
				break
			} else if mid+1 >= len(A) {
				return mid + 1
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}

	if n == len(A)-1 {
		return n + 1
	}

	l = n + 1
	r = len(A) - 1

	for l <= r {
		mid := (l + r) / 2

		if A[mid] > 0 && (A[mid-1] < 0 || A[mid-1] == 0) {
			p = mid
			break
		} else if A[mid] == 0 || A[mid] < 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	negative := n + 1
	postive := 0
	if p > -1 {
		postive = len(A) - p
	}

	if negative > postive {
		return negative
	}

	return postive
}
