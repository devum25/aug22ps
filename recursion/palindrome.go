package recursion

func Palindrome(s string) bool {
	return palindrome(s, 0, len(s)-1)
}

func palindrome(s string, i, j int) bool {
	if i > j {
		return true
	}
	if s[i] != s[j] {
		return false
	}
	i++
	j--
	return palindrome(s, i, j)
}
