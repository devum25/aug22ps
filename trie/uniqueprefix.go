package trie

// Given an array of words, Return an array of strings containing the smallest unique prefix for every word.
func GetUniquePrefix(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}
	res := make([]string, len(arr))
	t := Constructor()

	for i := 0; i < len(arr); i++ {
		t.Insert(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		res[i] = t.GetUniqueChar(arr[i])
	}

	return res
}
