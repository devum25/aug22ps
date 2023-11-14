package trie

func LongestCommonPrefix(arr []string) string {
	trie := NewTrieNode()

	for i := 0; i < len(arr); i++ {
		trie.Insert(arr[i])
	}

	return trie.GetLongestPrefix(arr[0], len(arr))
}

type TrieNode struct {
	hash map[byte]*TrieNode
	end  bool
	freq int
}

func NewTrieNode() TrieNode {
	return TrieNode{hash: make(map[byte]*TrieNode)}
}

func (t *TrieNode) Insert(word string) {
	curr := t

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			x := NewTrieNode()
			curr.hash[word[i]] = &x
		}
		curr = curr.hash[word[i]]
		curr.freq++
	}

	curr.end = true
}

func (t *TrieNode) GetLongestPrefix(word string, n int) string {
	curr := t
	var s string

	for i := 0; i < len(word); i++ {
		curr = curr.hash[word[i]]
		if curr.freq == n {
			s += string(word[i])
		} else if curr.freq < n {
			return s
		}
	}

	return s
}
