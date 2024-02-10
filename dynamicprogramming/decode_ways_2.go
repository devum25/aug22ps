package dynamicprogramming

func NumDecodings2(s string) int {
	dp := make([]int, len(s))

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return ways2(s, dp, len(s)-1)
}

func ways2(s string, dp []int, n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		if s[n] == '*' {
			return 9
		} else if s[n] != '0' {
			return 1
		}
		return 0
	}

	if n == 1 {
		if s[n] == '*' && s[n-1] == '*' {
			return 96
		}
		if s[n] == '0' && (s[n-1] == '1' || s[n-1] == '2') {
			return 1
		} else if s[n] != '0' && s[n-1] == '0' {
			return 0
		} else if s[n] == '*' {
			if s[n-1] == '1' {
				return 18
			} else if s[n-1] == '2' {
				return 15
			} else {
				return 9
			}
		} else if s[n] == '1' || s[n] == '2' {
			if s[n-1] == '*' {
				return 11
			}
		}
	}

	if dp[n] != -1 {
		return dp[n]
	}

	x := 0
	y := 0

	// check for single digit
	if s[n] != '0' {
		if s[n] == '*' {
			x = 9 * ways2(s, dp, n-1)
		} else {
			x = ways2(s, dp, n-1)
		}
	}

	if check(s, n-1, n) {
		if s[n] == '*' && s[n-1] == '*' {
			y = 15 * ways2(s, dp, n-2) // 11-19, 21-26
		} else if s[n] == '*' && s[n-1] == '1' {
			y = 18 * ways2(s, dp, n-2) // 11-19
		} else if s[n] == '*' && s[n-1] == '2' {
			y = 15 * ways2(s, dp, n-2) // 21-26
		} else {
			y = ways2(s, dp, n-2)
		}
	}

	dp[n] = x + y

	return dp[n]
}

func check(s string, i, j int) bool {
	if s[i] == '0' { // definetly can not be paired
		return false
	} else if s[i] == '*' {
		return true
	} else if s[i] == '1' {
		return true
	} else if s[i] == '2' {
		v := int(s[j])

		if v >= 48 && v <= 54 {
			return true
		}
		if s[j] == '*' {
			return true
		}
	}

	return false
}
