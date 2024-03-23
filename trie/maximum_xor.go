package trie

import (
	"math"
	"strconv"
)

func FindMaximumXOR(nums []int) int {
	strs := make([]string, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		s := convertToBinary(nums[i])
		if len(s) > max {
			max = len(s)
		}
		strs[i] = s
	}

	trie := NewTrieNode1()

	for i := 0; i < len(strs); i++ {
		strs[i] = trie.Insert(strs[i], max-len(strs[i]))
	}
	ans := 0
	for i := 0; i < len(strs); i++ {
		x := trie.SearchXOR(strs[i])
		x = x ^ nums[i]
		if x > ans {
			ans = x
		}
	}

	return ans
}

func convertToBinary(num int) string {
	var s string
	count := 32

	for num > 0 {
		temp := int(num % 2)
		s += strconv.Itoa(temp)
		num = num / 2
		count--
	}

	x := reverse(s)

	return x
}

func convertBinaryToInt(s string) int {
	ans := 0
	p := 0
	for i := len(s) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		ans += n * int(math.Pow(float64(2), float64(p)))
		p++
	}

	return ans
}

func reverse(s string) string {
	r := make([]byte, len(s))
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		r[j] = s[i]
		j++
	}

	return string(r)
}

func (t *TrieNode1) SearchXOR(s string) int {
	curr := t
	var res string
	ans := 0
	p := len(s) - 1
	for i := 0; i < len(s); i++ {
		x := '0'
		if s[i] == '0' {
			x = '1'
		}

		if _, ok := curr.hash[byte(x)]; ok {
			res += string(x)
			curr = curr.hash[byte(x)]
			n, _ := strconv.Atoi(string(x))
			ans += n * int(math.Pow(float64(2), float64(p)))
		} else {
			if x == '1' {
				res += string('0')
				curr = curr.hash[byte('0')]
				n, _ := strconv.Atoi(string('0'))
				ans += n * int(math.Pow(float64(2), float64(p)))
			} else {
				res += string('1')
				curr = curr.hash[byte('1')]
				n, _ := strconv.Atoi(string('1'))
				ans += n * int(math.Pow(float64(2), float64(p)))
			}
		}
		p--
	}

	return ans
}

type TrieNode1 struct {
	hash map[byte]*TrieNode1
}

func NewTrieNode1() TrieNode1 {
	return TrieNode1{hash: make(map[byte]*TrieNode1)}
}

func (t *TrieNode1) Insert(word string, diff int) string {
	curr := t
	i := 0
	c := 1
	var s string
	for i < len(word) || c <= diff {
		if c <= diff {
			x := NewTrieNode1()
			curr.hash['0'] = &x
			curr = curr.hash['0']
			c++
			s += string('0')
		} else if i < len(word) {
			if _, ok := curr.hash[word[i]]; !ok {
				x := NewTrieNode1()
				curr.hash[word[i]] = &x
			}
			curr = curr.hash[word[i]]
			s += string(word[i])
			i++

		}
	}

	return s
}
