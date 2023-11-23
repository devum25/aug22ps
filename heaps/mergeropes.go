package heaps

// Problem Description
// Given lengths of n ropes, the cost of merging two ropes is equal to sum of ropes getting merged
// Find minimum cost required to merge all the ropes in a single rope

func MergeRopes(arr []int) int {

	res := toMinHeap(arr)
	minHeap := NewMinHeap()
	minHeap.Items = res
	minHeap.size = len(res) - 1
	ans := 0
	s := 0
	e := len(arr) - 1

	for s < e {
		first := minHeap.DeleteMin()
		second := minHeap.DeleteMin()
		ans += first + second

		minHeap.Insert(first + second)
		s++
	}

	return ans

}

func toMinHeap(arr []int) []int {
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
			min := i
			if l < len(arr) && res[l] < res[min] {
				min = l
			}
			if r < len(arr) && res[r] < res[min] {
				min = r
			}
			if i == min {
				break
			}
			res[min], res[i] = res[i], res[min]
			i = min
		}
		s--
	}

	return res

}
