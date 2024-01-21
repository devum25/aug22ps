package dynamicprogramming

// Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
// An interleaving of two strings s and t is a configuration where s and t are divided into n and m
// substrings
//  respectively, such that:

// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
// Note: a + b is the concatenation of strings a and b.

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if s3 == "" && s1 == "" && s2 == "" {
		return true
	} else if s3 == "" && (s1 != "" || s2 != "") {
		return false
	}
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return interleave(s1, s2, s3, len(s1), len(s2), len(s3), dp)
}

func interleave(s1, s2, s3 string, i, j, n int, dp [][]int) bool {
	if i == 0 && n == 0 && j == 0 {
		return true
	} else if n == 0 {
		return false
	}
	if dp[i][j] == -1 {
		if i > 0 && s1[i-1] == s3[n-1] && interleave(s1, s2, s3, i-1, j, n-1, dp) {
			dp[i][j] = 1
			return true
		} else if j > 0 && s2[j-1] == s3[n-1] && interleave(s1, s2, s3, i, j-1, n-1, dp) {
			dp[i][j] = 1
			return true
		} else {
			dp[i][j] = 0
		}
	}

	return dp[i][j] == 1
}
