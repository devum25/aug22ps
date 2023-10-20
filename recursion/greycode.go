package recursion

import "math"

func GenerateGreyCode(n int) []int {
	s := make([]string, 0)
	s = helperGray(n, s)
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		k := 0
		val := 0
		for j := len(s[i]) - 1; j >= 0; j-- {
			if s[i][j] == '0' {
				val = val + 0
			} else {
				val = val + 1*int(math.Pow(float64(2), float64(k)))
			}
			k++
		}
		res[i] = val
	}

	return res
}

func helperGray(n int, s []string) []string {
	if n == 1 {
		s = append(s, "0", "1")
		return s
	}

	s = helperGray(n-1, s)
	x := make([]string, 0)
	k := len(s)
	for i := 0; i < k; i++ {
		x = append(x, "0"+s[i])
	}

	for i := k - 1; i >= 0; i-- {
		x = append(x, "1"+s[i])
	}

	return x
}
