package strings

import "fmt"

func RomanToInt(s string) int {
	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	x := toInt(s, len(s)-1, "", m)
	fmt.Print(x)
	return x
}

func toInt(str string, i int, prev string, hash map[string]int) int {
	ans := 0
	if i < 0 {
		return 0
	}

	if prev == "" {
		ans = toInt(str, i-1, string(str[i]), hash) + hash[string(str[i])]
	} else {
		if hash[prev] > hash[string(str[i])] {
			ans = toInt(str, i-1, string(str[i]), hash) - hash[string(str[i])]
		} else {
			ans = toInt(str, i-1, string(str[i]), hash) + hash[string(str[i])]
		}
	}

	return ans
}
