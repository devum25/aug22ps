package trie

import (
	"sort"
)

// LEETCODE CONTEST 388
// 3076. Shortest Uncommon Substring in an Array
// You are given an array arr of size n consisting of non-empty strings.

// Find a string array answer of size n such that:

// answer[i] is the shortest substring of arr[i] that does not occur as a substring in any other string in arr. If multiple such substrings exist, answer[i] should be the lexicographically smallest. And if no such substring exists, answer[i] should be an empty string.
// Return the array answer.
// Input: arr = ["cab","ad","bad","c"]
// Output: ["ab","","ba",""]
func ShortestSubstring(arr []string) []string {

	hash := make(map[int]map[string]bool)
	res := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		hash[i] = populatesubstrings(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		res[i] = getUniqueSubstring(i, len(arr), arr[i], hash)
	}

	return res
}

func populatesubstrings(s string) map[string]bool {
	lst := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			lst[s[i:j+1]] = true
		}
	}

	return lst
}

func getUniqueSubstring(idx, n int, str string, hash map[int]map[string]bool) string {
	lst := hash[idx]

	size := len(str)
	arr := make([]string, 0)
	for sub := range lst {
		found := false
		x := 0
		for x < n {
			if x == idx {
				x++
				continue
			}
			temp := hash[x]
			if temp[sub] {
				x++
				found = true
				break
			}
			x++
		}
		if !found {
			if len(arr) == 0 {
				arr = append(arr, sub)
				size = len(sub)
			} else {
				if len(sub) <= size {
					if len(sub) < size {
						arr = arr[len(arr):]
						size = len(sub)
					}
					arr = append(arr, sub)
				}
			}
		}
	}

	if len(arr) == 0 {
		return ""
	}

	sort.Strings(arr)

	return arr[0]
}
