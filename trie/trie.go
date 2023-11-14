package trie

type Trie struct {
	hash map[byte]*Trie
	end  bool
	freq int
}

func Constructor() Trie {
	return Trie{hash: make(map[byte]*Trie)}
}

func (this *Trie) Insert(word string) {
	curr := this

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			x := Constructor()
			curr.hash[word[i]] = &x
		}
		curr = curr.hash[word[i]]
		curr.freq++
	}
	curr.end = true
}

func (this *Trie) Search(word string) bool {
	curr := this

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			return false
		}
		curr = curr.hash[word[i]]
	}

	return curr.end
}

func (this *Trie) GetUniqueChar(word string) string {
	curr := this
	var s string
	for i := 0; i < len(word); i++ {
		curr = curr.hash[word[i]]
		if curr.freq > 1 {
			s += string(word[i])
		} else {
			s += string(word[i])
			return s
		}
	}

	return s
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this

	for i := 0; i < len(prefix); i++ {
		if _, ok := curr.hash[prefix[i]]; !ok {
			return false
		}
		curr = curr.hash[prefix[i]]
	}

	return true
}
