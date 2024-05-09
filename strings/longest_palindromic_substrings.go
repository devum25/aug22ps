package strings

import "math"

func LongestPalindrome(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	}
	ans := math.MinInt
	ansi := -1
	ansj := -1
	for i := 0; i < len(s); i++ {
		odd, x, y := getPalindromeLength(s, i, i)
		if odd > ans {
			if (x+1) >= 0 && (y-1) < len(s) {
				ansi = x + 1
				ansj = y - 1
				ans = odd
			}
		}
		even, k, m := getPalindromeLength(s, i, i+1)
		if even > ans {
			if (k+1) >= 0 && (m-1) < len(s) {
				ansi = k + 1
				ansj = m - 1
				ans = even
			}
		}
	}
	return s[ansi : ansj+1]
}

func getPalindromeLength(s string, i, j int) (int, int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1, i, j
}
