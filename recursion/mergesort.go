package recursion

func SortArray(n []int) []int {
	return mergesort(n)
}

func mergesort(n []int) []int {
	// res := make([]int, 0)
	if len(n) <= 1 {
		return n
	}

	a := mergesort(n[0 : len(n)/2]) //divide and conquer
	b := mergesort(n[len(n)/2:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	res := make([]int, 0)

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
			continue
		}
		res = append(res, b[j])
		j++
	}

	if i < len(a) {
		for i < len(a) {
			res = append(res, a[i])
			i++
		}
	}
	if j < len(b) {
		for j < len(b) {
			res = append(res, b[j])
			j++
		}
	}

	return res
}
