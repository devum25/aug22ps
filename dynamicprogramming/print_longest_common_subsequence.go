package dynamicprogramming

func PrintLcs(s1, s2 string) string {
	dp := longestCommonSubsequenceMatrix(s1, s2)
	res := ""
	row := len(dp) - 1
	col := len(dp[0]) - 1

	for row >= 1 && col >= 1 {
		if s1[row-1] == s2[col-1] {
			res += string(s1[row-1])
			row--
			col--
		} else {
			if dp[row-1][col] > dp[row][col-1] {
				row--
			} else {
				col--
			}
		}
	}
	x := reverse(res)
	return x
}

func reverse(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func longestCommonSubsequenceMatrix(text1 string, text2 string) [][]int {
	dp := make([][]int, len(text1)+1)
	n := len(text2) + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp
}
