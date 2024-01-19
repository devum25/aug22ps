package dynamicprogramming

func IterativeRegularExpression(s, p string) bool {
	dp := make([][]int, len(s)+1)
	n := len(p) + 1

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	r := 0
	c := 1
	for c < len(dp[r]) {
		if p[c] == '*' {
			dp[0][c] = dp[0][c-1]
		}
		c++
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				x := dp[i][j-2]
				if x == 1 {
					dp[i][j] = x
					continue
				} else {
					if p[j-1] == s[i-1] || p[j-1] == '.' {
						dp[i][j] = dp[i-1][j]
					}
				}
			}
		}
	}

	x := dp[len(s)][len(p)]
	return x == 1
}
