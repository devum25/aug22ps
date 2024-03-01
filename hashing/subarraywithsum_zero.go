package hashset

// Given an array, return true if there exists a subarray with sum = 0
// Create pf sum
// if values are repeating then that means subarray between them is cancelling out the contribution.
func SubarrayWithSumZero(A []int) bool {
	pf := make([]int, len(A))
	pf[0] = A[0]
	hash := make(map[int]int)
	for i := 1; i < len(A); i++ {
		pf[i] = pf[i-1] + A[i]
		if pf[i] == 0 {
			return true
		}
		hash[pf[i]]++
	}

	for _, v := range hash {
		if v > 1 {
			return true
		}
	}

	return false
}
