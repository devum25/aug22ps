package trie

func SpellChecker(dictionary []string, words []string) []int {
	res := make([]int, len(words))
	trie := NewTrieNodeDic()
	for i := 0; i < len(dictionary); i++ {
		trie.Insert(dictionary[i])
	}

	for i := 0; i < len(words); i++ {
		if trie.Search(words[i]) {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}

	return res
}

type TrieNodeDic struct {
	hash map[byte]*TrieNodeDic
	end  bool
}

func NewTrieNodeDic() TrieNodeDic {
	return TrieNodeDic{hash: make(map[byte]*TrieNodeDic)}
}

func (t *TrieNodeDic) Insert(word string) {
	curr := t

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			x := NewTrieNodeDic()
			curr.hash[word[i]] = &x
		}
		curr = curr.hash[word[i]]
	}
	curr.end = true
}

func (t *TrieNodeDic) Search(word string) bool {
	curr := t

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			return false
		}
		curr = curr.hash[word[i]]
	}

	return curr.end
}
