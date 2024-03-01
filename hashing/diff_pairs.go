package hashset

// Given an arr A, and an integer k.
// Return true if there exists a pair A[i]-A[j] = k ; i < j

func CheckDiff(A []int, k int) bool {

	hash := make(map[int]bool)
	hash[A[0]] = true

	for i := 1; i < len(A); i++ {
		if hash[k+A[i]] {
			return true
		}
		hash[k+A[i]] = true
	}

	return false
}
