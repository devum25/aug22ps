package strings

func Palindrome(s string) bool {

	i := 0
	j := len(s) - 1

	for i < j {

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
