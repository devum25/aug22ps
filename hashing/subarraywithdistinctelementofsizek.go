package hashset

func SubArrayWithDistinctElementOfSizeK(A []int, k int) []int {
	res := make([]int, 0)

	hash := make(map[int]int)
	count := 0
	for i := 0; i < k; i++ {
		if _, ok := hash[A[i]]; ok {
			hash[A[i]]++
			continue
		}
		hash[A[i]]++
		count++
	}

	res = append(res, count)

	s := 1
	e := k
	for e < len(A) {
		hash[A[s-1]]--
		if hash[A[s-1]] == 0 {
			delete(hash, A[s-1])
			count--
		}
		if _, ok := hash[A[e]]; ok {
			hash[A[e]]++
			res = append(res, count)
		} else {
			hash[A[e]]++
			count++
			res = append(res, count)
		}
		s++
		e++
	}

	return res

}
