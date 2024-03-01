package hashset

//Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

func NonRepeatingCharacter(s string) int {
	hash := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			return i
		}
	}

	return -1
}
