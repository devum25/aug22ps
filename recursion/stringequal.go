package recursion

func Compare(s string, t string, i int) bool {
	if len(s) != len(t) {
		return false
	}

	if i == len(s) {
		return true
	}

	if s[i] != t[i] {
		return false
	}

	return Compare(s[i+1:], t[i+1:], i)
}
