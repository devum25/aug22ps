package recursion

func GenerateParanthesis(n int) []string {
	result := []string{}
	solve(n, n, &result, []byte{})
	return result
}

func solve(open, close int, res *[]string, sub []byte) {
	if open == 0 && close == 0 {
		*res = append(*res, string(sub))
		return
	}

	if open > 0 {
		solve(open-1, close, res, append(sub, '('))
	}

	if open < close {
		solve(open, close-1, res, append(sub, ')'))
	}
}
