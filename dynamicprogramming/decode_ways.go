package dynamicprogramming

// 91. Decode Ways

// A message containing letters from A-Z can be encoded into numbers using the following mapping:
// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
// To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:
// "AAJF" with the grouping (1 1 10 6)
// "KJF" with the grouping (11 10 6)
// Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
// Given a string s containing only digits, return the number of ways to decode it.
// The test cases are generated so that the answer fits in a 32-bit integer.

// Example 1:
// Input: s = "12"
// Output: 2
// Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
// Example 2:
// Input: s = "226"
// Output: 3
// Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
// Example 3:
// Input: s = "06"
// Output: 0
// Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

// func NumDecodings(s string) int {
// 	dp := make([]int, len(s)+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = -1
// 	}
// 	x := decode(s, &dp, len(s)-1)
// 	return x
// }

// func decode(n string, dp *[]int, i int) int {
// 	if i == 0 {
// 		if n[i] != '0' {
// 			return 1
// 		}
// 		return 0
// 	}

// 	if i == 1 {
// 		temp := 0
// 		if n[i] != '0' && n[i-1] != '0' {
// 			temp++
// 		}
// 		if check(n, i-1, i) {
// 			temp++
// 			return temp
// 		}
// 		return temp
// 	}

// 	if (*dp)[i] != -1 {
// 		return (*dp)[i]
// 	}
// 	x := 0
// 	y := 0
// 	if n[i] != '0' {
// 		x = decode(n, dp, i-1) //single digit
// 	}
// 	if check(n, i-1, i) {
// 		y = decode(n, dp, i-2) //double digit
// 	}

// 	(*dp)[i] = x + y
// 	return (*dp)[i]
// }

// func check(s string, i, j int) bool {
// 	if s[i] == '0' {
// 		return false
// 	} else if s[i] == '2' {
// 		return checkval(s[j])
// 	} else if s[i] == '1' {
// 		return true
// 	}
// 	return false
// }

// func checkval(c byte) bool {
// 	s := string(c)

// 	v, _ := strconv.Atoi(s)
// 	return v <= 6
// }

///////////////////////////////////////////////////////
func NumDecodings(s string) int {
	dp := make([]int, len(s))

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return ways1(s, dp, len(s)-1)
}

func ways1(s string, dp []int, n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		if s[n] != '0' {
			return 1
		}
		return 0
	}
	if n == 1 {
		temp := 0
		if s[n-1] != '0' {
			temp++
		}
		if checkForPair(s, n-1, n) {
			temp++
		}
		return temp
	}

	if dp[n] != -1 {
		return dp[n]
	}

	x, y := 0, 0

	if s[n] != '0' {
		x = ways1(s, dp, n-1)
	}
	if checkForPair(s, n-1, n) {
		y = ways1(s, dp, n-2)
	}

	dp[n] = x + y

	return dp[n]
}

func checkForPair(s string, i, j int) bool {
	if s[i] == '0' {
		return false
	} else if s[i] == '1' {
		return true
	} else if s[i] == '2' {
		v := int(s[j])

		if v >= 48 && v <= 54 {
			return true
		}
	}

	return false
}
