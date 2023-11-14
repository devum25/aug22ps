package arrayquestion

func MaxPostiveSubArray(arr []int) []int {
	ans := 0
	s := 0
	e := 0
	res := make([]int, 0)
	prev := 0
	for e < len(arr) {
		if arr[e] > 0 {
			ans = max(ans, e-s+1)
		} else {
			if ans > prev {
				res = make([]int, 0)
				for i := s; i < e; i++ {
					res = append(res, arr[i])
				}
			}
			prev = ans
			s = e + 1
		}
		e++
	}
	if ans > prev {
		res = make([]int, 0)
		for i := s; i < e; i++ {
			res = append(res, arr[i])
		}
	}
	return res
}
