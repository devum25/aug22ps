package recursion

import (
	"strconv"
	"strings"
)

func KthGrammar(n int, k int) int {
	x := make([]string, 0)
	x = helper(n, n, x)
	s := x[n-1]

	arr := strings.Split(s, "")
	e, _ := strconv.Atoi(arr[k])
	return e
}

func helper(n, k int, s []string) []string {
	if k == 0 {
		return s
	}
	if k == n {
		s = append(s, string('0'))
		k--
		return helper(n, k, s)

	}
	res := make([]byte, 0)
	val := n - (k + 1)
	x := s[val]
	for i := 0; i < len(x); i++ {
		if x[i] == '0' {
			res = append(res, '0', '1')
		} else if x[i] == '1' {
			res = append(res, '1', '0')
		}
	}

	s = append(s, string(res))
	k--
	return helper(n, k, s)
}
