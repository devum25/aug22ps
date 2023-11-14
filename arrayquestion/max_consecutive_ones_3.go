package arrayquestion

// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8
// 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1
func MaxConsecutiveOneWithKFlips(arr []int, k int) int {
	ans := 0
	s := 0
	e := 0
	zeros := 0
	for e < len(arr) {
		if arr[e] == 1 {
			ans = max(ans, e-s+1)
		} else {
			zeros++
			for zeros > k {
				if arr[s] == 0 {
					zeros--
				}
				s++
			}
		}
		ans = max(ans, e-s+1)
		e++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
