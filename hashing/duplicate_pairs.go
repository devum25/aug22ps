package hashset

// Given an array of N elements, count the number of duplicate pairs (i,j) i.e A[i] == A[j]; i != j

// NC2 = N*(N-1)/2

func DuplicatePairs(A []int) int {
	ans := 0

	hash := make(map[int][]int)

	for i := 0; i < len(A); i++ {
		hash[A[i]] = append(hash[A[i]], i)
	}

	for _, v := range hash {
		ans += ((len(v)) * (len(v) - 1)) / 2
	}

	return ans
}

// Another Approach

func DuplicatePairs1(A []int) int {
	ans := 0

	hash := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := hash[A[i]]; ok {
			ans += hash[A[i]]
			hash[A[i]]++
		} else {
			hash[A[i]]++
		}
	}

	return ans
}
