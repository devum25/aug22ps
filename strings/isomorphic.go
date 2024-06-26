package strings

// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

func IsIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hash := make(map[byte]byte)
	taken := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := hash[t[i]]; !ok {
			if !taken[s[i]] {
				hash[t[i]] = s[i]
				taken[s[i]] = true
			} else {
				return false
			}
		} else {
			if hash[t[i]] != s[i] {
				return false
			}
		}

	}

	return true

}
