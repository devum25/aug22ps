package hashset

func NumberOfSpecialChars(word string) int {
	hash := make(map[byte]bool)
	count := 0
	visited := make(map[byte]bool)

	for i := 0; i < len(word); i++ {

		if int(word[i]) >= 65 && int(word[i]) <= 91 {
			x := word[i] - 'A'
			if _, ok := hash[x]; !ok {
				hash[x] = true
			}
			if hash[byte(x+'a')] && !visited[word[i]] {
				count++
				visited[x+'a'] = true
				visited[word[i]] = true
			}
		}
		if int(word[i]) >= 97 && int(word[i]) <= 123 {
			x := word[i] - 'a'
			if _, ok := hash[word[i]]; !ok {
				hash[word[i]] = true
			}
			if hash[byte(x+'A')] && !visited[word[i]] {
				count++
				visited[x+'A'] = true
				visited[word[i]] = true
			}
		}
	}

	return count

}
