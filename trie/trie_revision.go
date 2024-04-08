package trie

type MyTrie struct {
	hash map[byte]*MyTrie
	end  bool
}

func NewMyTrie() *MyTrie {
	return &MyTrie{hash: make(map[byte]*MyTrie)}
}

func (t *MyTrie) Insert(str string) {
	curr := t

	for i := 0; i < len(str); i++ {
		if _, ok := curr.hash[str[i]]; !ok {
			trie := NewMyTrie()
			curr.hash[str[i]] = trie
		}
		curr = curr.hash[str[i]]
	}

	curr.end = true
}

func (t *MyTrie) Search(str string) bool {
	curr := t

	for i := 0; i < len(str); i++ {
		if _, ok := curr.hash[str[i]]; !ok {
			return false
		}
		curr = curr.hash[str[i]]
	}

	return curr.end
}
