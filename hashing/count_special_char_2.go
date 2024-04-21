package hashset

func NumberOfSpecialChars2(word string) int {
	hash := make(map[byte]bool)
	visited := make(map[byte]bool)
	count := 0
	for i := 0; i < len(word); i++ {

		if int(word[i]) >= 97 && int(word[i]) <= 123 {
			//x := word[i]-'a'
			if visited[word[i]] {
				count--
				hash[word[i]] = false
				visited[word[i]] = false
				continue
			}

			if _, ok := hash[word[i]]; !ok {
				x := word[i] - 'a'
				if !hash[byte(x+'A')] {
					hash[word[i]] = true
				}
			}
			continue

		}

		if int(word[i]) >= 65 && int(word[i]) <= 91 {
			x := word[i] - 'A'
			if _, ok := hash[word[i]]; !ok {
				// if hash[byte(x+'a')] {
				hash[word[i]] = true
				// }
			}

			if hash[byte(x+'a')] && !visited[byte(x+'a')] {
				count++
				visited[byte(x+'a')] = true
			}
		}

	}

	return count
}
