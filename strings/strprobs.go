package strings

import "math"

func ToggleChars(str string) string {
	str = "AbhJkLmmOP"

	dif := 'a' - 'A'

	newStr := make([]byte, len(str))

	for i := 0; i < len(str); i++ {

		if str[i] >= 'A' && str[i] <= 'Z' {
			// str[i] = str[i]+byte(dif)
			val := str[i] + byte(dif)
			newStr[i] = val
		} else if str[i] >= 'a' && str[i] <= 'z' {
			val := str[i] - byte(dif)
			newStr[i] = val
		}
	}

	return string(newStr)
}

func ArrangeInLexicographical(str string) string { //Bubblesort //O(N2)
	str = "zunmabffp" //aabbcdde

	newString := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		newString[i] = str[i]
	}

	for i := 0; i < len(newString); i++ {
		for j := 0; j < len(newString)-1; j++ {
			if newString[i] < newString[j] {
				newString[i], newString[j] = newString[j], newString[i]
			}
		}
	}

	return string(newString)
}

func ArrangeInLexicographicalOptimal(str string) string {
	arr := make([]int, 26)

	str = "deabadcb"
	x := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		arr[str[i]-'a']++
	}
	j := 0
	for i := 0; i < len(arr); i++ {
		k := arr[i]
		for k > 0 {

			x[j] = byte(i) + 'a'
			k--
			j++
		}
	}

	return string(x)
}

// idea is to consider ith element as center and travel left and right of ith element to check if it's palindrome or not.
// ith element can form odd palindrome or even palindrome.
// check for both odd and even and get max

func GetLargetSubStringPalindrome(str string) int { // get largest
	str = "xbdyzzydbdyzydx"
	max := 1
	for i := 0; i < len(str); i++ {
		odd := getLenghtOfPalindrome(str, i, i)
		even := getLenghtOfPalindrome(str, i, i+1)
		x := math.Max(float64(odd), float64(even))
		max = int(math.Max(float64(max), float64(x)))
	}

	return max
}

func getLenghtOfPalindrome(str string, i, j int) int {

	for i >= 0 && j < len(str) && str[i] == str[j] {
		i--
		j++
	}

	return j - i - 1
}
