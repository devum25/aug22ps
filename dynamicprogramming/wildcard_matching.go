package dynamicprogramming

// 44. Wildcard Matching
// Hard
// 8K
// 324
// Companies
// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).

func IsMatch(s string, p string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return tobool(match(s, p, len(s), len(p), dp))
}

func match(s, p string, ls, lp int, dp [][]int) int {
	if ls == 0 && lp == 0 {
		return 1
	}
	if lp <= 0 {
		return 0
	}
	if ls <= 0 {
		if p[lp-1] == '*' {
			return match(s, p, ls, lp-1, dp)
		} else {
			return 0
		}
	}
	if dp[ls][lp] != -1 {
		return dp[ls][lp]
	}

	if isChar(p[lp-1]) {
		if s[ls-1] == p[lp-1] {
			dp[ls][lp] = match(s, p, ls-1, lp-1, dp)
		}
	}
	if p[lp-1] == '?' {
		dp[ls][lp] = match(s, p, ls-1, lp-1, dp)
	}
	if p[lp-1] == '*' {
		x := match(s, p, ls-1, lp, dp)
		y := match(s, p, ls, lp-1, dp)

		if x == 1 || y == 1 {
			dp[ls][lp] = 1
		} else {
			dp[ls][lp] = 0
		}
	}

	return dp[ls][lp]
}

func isChar(c byte) bool {
	s := int('a')
	e := int('z')
	i := int(c)
	if i >= s && i <= e {
		return true
	}

	return false
}

func tobool(x int) bool {
	return x == 1
}
