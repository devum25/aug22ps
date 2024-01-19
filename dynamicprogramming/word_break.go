package dynamicprogramming

// 139. Word Break
// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:
// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

func WordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)

	for i := 0; i < len(wordDict); i++ {
		hash[wordDict[i]] = true
	}

	dp := make([]int, len(s)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return checkWord(s, 0, hash, dp)
}

func checkWord(str string, s int, hash map[string]bool, dp []int) bool {
	if dp[s] == -1 {
		if s == len(str) {
			dp[s] = 1
		} else {
			for i := s; i < len(str); i++ {
				if hash[str[s:i+1]] && checkWord(str, i+1, hash, dp) {
					dp[s] = 1
				}
			}
			if dp[s] == -1 {
				dp[s] = 0
			}
		}
	}

	return dp[s] == 1
}
