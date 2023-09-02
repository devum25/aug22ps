package hashset

import "fmt"

func Distinctelemntofsizek() []int {
	res := make([]int, 0)

	// A := []int{6, 3, 7, 3, 8, 6, 9}
	// k:=3
	A := []int{1, 1, 2, 2}
	k := 2

	count := 0

	for i := 0; i < len(A) && (len(A)-i >= k); i++ {
		x := make(map[int]int)
		count = 0
		j := i
		c := k
		for j < len(A) && c > 0 {
			if _, ok := x[A[j]]; !ok {
				fmt.Println(A[j])
				x[A[j]]++
				count++
			}
			c--
			j++
		}

		fmt.Println(count)
		res = append(res, count)
	}

	return res
}

func DistinctelemntofsizekON() []int {
	res := make([]int, 0)

	// A := []int{6, 3, 7, 3, 8, 6, 9}
	// k := 3

	A := []int{1, 1, 2, 2}
	k := 2

	kv := make(map[int]int)

	for i := 0; i < k; i++ {
		kv[A[i]]++
	}

	count := len(kv)
	res = append(res, count)

	s := 1
	e := k

	for e < len(A) {
		kv[A[s-1]]--
		if val, ok := kv[A[s-1]]; ok {
			if val == 0 {
				delete(kv, A[s-1])
			}
		}

		kv[A[e]]++
		res = append(res, len(kv))
		s++
		e++
	}

	return res
}
