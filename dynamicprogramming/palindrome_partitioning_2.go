package dynamicprogramming

import "math"

// 132. Palindrome Partitioning II
// Given a string s, partition s such that every substring of the partition is a palindrome
// Return the minimum cuts needed for a palindrome partitioning of s.

// Example 1:
// Input: s = "aab"
// Output: 1
// Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

func MinCut(s string) int {
	dp := make([]int, len(s)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return cut(s, dp, 0)
}

func cut(str string, dp []int, s int) int {
	if dp[s] == -1 {
		if s == len(str) {
			dp[s] = 0
		} else {
			ans := math.MaxInt
			for i := s; i < len(str); i++ {
				if isValidPalindrome(str[s : i+1]) {
					x := 1 + cut(str, dp, i+1)
					ans = min(ans, x)
				}
			}
			if ans != math.MaxInt {
				dp[s] = ans
			}
		}
	}

	return dp[s]
}

func isValidPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] == s[j] {
			return true
		}
		i++
		j--
	}
	return false
}
