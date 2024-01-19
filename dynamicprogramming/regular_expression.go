package dynamicprogramming

// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

func IsRegularExpression(s, p string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	x := ismatch(s, p, len(s), len(p), dp)
	return x == 1
}

func ismatch(s, p string, i, j int, dp [][]int) int {
	if dp[i][j] == -1 {
		if i == 0 && j == 0 {
			dp[i][j] = 1
		} else if j == 0 {
			dp[i][j] = 0
		} else if i == 0 {
			dp[i][j] = ismatch(s, p, i, j-1, dp)
		} else if isChar(p[j-1]) {
			if s[i-1] == p[j-1] {
				dp[i][j] = ismatch(s, p, i-1, j-1, dp)
			}
		} else if p[j-1] == '.' {
			dp[i][j] = ismatch(s, p, i-1, j-1, dp)
		} else if p[j-1] == '*' {
			x := ismatch(s, p, i, j-2, dp)
			y := ismatch(s, p, i-1, j, dp)
			if x == 1 || y == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}

	}

	return dp[i][j]
}
