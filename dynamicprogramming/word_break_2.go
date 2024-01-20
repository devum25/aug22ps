package dynamicprogramming

// Given a string s and a dictionary of strings wordDict, add spaces in s
// to construct a sentence where each word is a valid dictionary word.
// Return all such possible sentences in any order.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:
// Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// Output: ["cats and dog","cat sand dog"]

func WordBreak2(s string, wordDict []string) []string {
	memo := make(map[int][]string, len(s))

	hash := make(map[string]bool)

	for i := 0; i < len(wordDict); i++ {
		hash[wordDict[i]] = true
	}

	x := sentences(s, hash, 0, memo)
	return x
}

func sentences(str string, hash map[string]bool, s int, dp map[int][]string) []string {
	if s == len(str) {
		return []string{""}
	}

	if r, ok := dp[s]; ok {
		return r
	}
	ans := make([]string, 0)
	for i := s; i < len(str); i++ {
		if hash[str[s:i+1]] {
			next := sentences(str, hash, i+1, dp)
			temp := str[s : i+1]
			for _, n := range next {
				if len(n) > 0 {
					tempWithSpace := temp + " " + n
					ans = append(ans, tempWithSpace)
				} else {
					ans = append(ans, temp)
				}
			}
		}
	}

	dp[s] = ans

	return ans
}
