package dynamicprogramming

func WordBreakIterative(s string, words []string) bool {
	hash := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		hash[words[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s); i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if hash[s[i:j+1]] && dp[j+1] {
				dp[i] = true
			}
		}
	}

	return dp[0]
}
