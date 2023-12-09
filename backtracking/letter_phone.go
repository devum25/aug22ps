package backtracking

import "strings"

// Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given below.
//  Note that 1 does not map to any letters.
// 2-abc
// 3-def
// 4-ghi
// 5-jkl
// 6-mno
// 7-pqrs
// 8-tuv
// 9-wxyz

// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

func LetterCombinations(digits string) []string {
	hash := make(map[string]string)

	hash["2"] = "abc"
	hash["3"] = "def"
	hash["4"] = "ghi"
	hash["5"] = "jkl"
	hash["6"] = "mno"
	hash["7"] = "pqrs"
	hash["8"] = "tuv"
	hash["9"] = "wxyz"

	// arr := make([]int, len(digits))
	input := make([][]string, 0)

	for i := 0; i < len(digits); i++ {
		x := make([]string, 0)
		str := hash[string(digits[i])]
		for j := 0; j < len(str); j++ {
			x = append(x, string(str[j]))
		}
		input = append(input, x)
	}

	ans := make([]string, 0)
	curr := make([]string, 0)
	var s string
	getCombination(input, 0, 0, &s, &curr, &ans)

	return ans
}

func getCombination(arr [][]string, idx int, arrIdx int, str *string, curr *[]string, ans *[]string) {
	if idx == len(arr) {
		temp := make([]string, len(*curr))
		for i := 0; i < len(temp); i++ {
			temp[i] = (*curr)[i]
		}
		s := strings.Join(temp, "")
		*ans = append(*ans, s)
		return
	}

	for i := 0; i < len(arr[idx]); i++ {
		*curr = append(*curr, arr[idx][i])
		getCombination(arr, idx+1, arrIdx+1, str, curr, ans)
		*curr = (*curr)[:len(*curr)-1]
	}
}
