package trie

// 676. Implement Magic Dictionary

// Design a data structure that is initialized with a list of different words. Provided a string, you should determine if you can change exactly one character in this string to match any word in the data structure.

// Implement the MagicDictionary class:

// MagicDictionary() Initializes the object.
// void buildDict(String[] dictionary) Sets the data structure with an array of distinct strings dictionary.
// bool search(String searchWord) Returns true if you can change exactly one character in searchWord to match any string in the data structure, otherwise returns false.

// Input
// ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
// [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
// Output
// [null, null, false, true, false, false]

type MagicDictionary struct {
	hash map[byte]*MagicDictionary
	end  bool
}

func ConstructorMagicDictionary() MagicDictionary {
	return MagicDictionary{hash: make(map[byte]*MagicDictionary)}
}

// func (this *MagicDictionary) BuildDict(dictionary []string) {
// 	curr := this

// 	for i := 0; i < len(dictionary); i++ {
// 		curr.Insert(dictionary[i])
// 	}
// }

// func (this *MagicDictionary) Insert(word string) {
// 	curr := this

// 	for i := 0; i < len(word); i++ {
// 		if _, ok := curr.hash[word[i]]; !ok {
// 			x := ConstructorMagicDictionary()
// 			curr.hash[word[i]] = &x
// 		}
// 		curr = curr.hash[word[i]]
// 	}
// 	curr.end = true
// }

// func (this *MagicDictionary) Search(word string) bool {
// 	curr := this
// 	return search(curr, 0, word, 0)
// }

// func search(node *MagicDictionary, idx int, word string, flag int) bool {

// 	curr := node

// 	// for i := 0; i < len(word); i++ {
// 	// 	if _, ok := curr.hash[word[i]]; !ok {

// 	// 	}
// 	// 	curr = curr.hash[word[i]]
// 	// }

// 	search(curr.hash[word[idx]], idx+1, word, flag)

// 	return false
// }

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		m.insert(word)
	}
}

func (m *MagicDictionary) insert(word string) {
	curr := m

	for i := 0; i < len(word); i++ {
		if _, ok := curr.hash[word[i]]; !ok {
			x := ConstructorMagicDictionary()
			curr.hash[word[i]] = &x
		}
		curr = curr.hash[word[i]]
	}
	curr.end = true
}

func (m *MagicDictionary) Search(searchWord string) bool {
	return dfs(m, searchWord, 0, 1)
}

func dfs(r *MagicDictionary, w string, i int, limit int) bool {
	// base case
	if limit < 0 {
		return false
	}
	if i == len(w) {
		return r.end && limit == 0
	}

	// iterate current node's all children
	ch := w[i]
	for c, t := range r.hash {
		if t == nil {
			continue
		}
		// c == ch, represent don't need change.
		if c == ch && dfs(t, w, i+1, limit) {
			return true
		}
		// c != ch, represent consume one chance for change
		if c != ch && dfs(t, w, i+1, limit-1) {
			return true
		}
	}
	return false
}
