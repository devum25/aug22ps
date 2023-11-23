package heaps

func CreateMaxHeap(arr []int) []int {
	res := make([]int, len(arr))
	n := len(arr)
	leaf := (n + 1) / 2
	for i := leaf; i < len(arr); i++ {
		res[i] = arr[i]
	}
	s := leaf - 1
	for s >= 0 {
		res[s] = arr[s]
		i := s
		for i < len(arr) {
			l := 2*i + 1
			r := 2*i + 2
			max := i
			if l < len(arr) && res[l] > res[max] {
				max = l
			}
			if r < len(arr) && res[r] > res[max] {
				max = r
			}
			if i == max {
				break
			}
			res[max], res[i] = res[i], res[max]
			i = max
		}

		s--
	}

	return res
}
