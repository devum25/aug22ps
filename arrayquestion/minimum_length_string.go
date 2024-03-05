package arrayquestion

func MinimumLength(s string) int {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return j - i + 1
		}

		i = i + 1
		for i < j && s[i] == s[i-1] {
			i++
		}
		j = j - 1
		for j > i && s[j] == s[j+1] {
			j--
		}
	}

	return j - i + 1
}
