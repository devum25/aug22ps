package strings

import "strings"

func CapitalizeTitle(title string) string {
	arr := strings.Split(title, " ")

	for j := 0; j < len(arr); j++ {
		if len(arr) <= 2 {
			s := arr[j]
			temp := make([]byte, len(s))
			for i := 0; i < len(s); i++ {
				if s[i] >= 65 && s[i] <= 90 {
					d := s[i] - 'A'
					temp[i] = d + 'a'
				} else {
					temp[i] = s[i]
				}
			}
			arr[j] = string(temp)
		} else {
			s := arr[j]
			temp := make([]byte, len(s))
			for i := 0; i < len(s); i++ {
				if s[i] >= 65 && s[i] <= 90 {
					if i > 0 {
						d := s[i] - 'A'
						temp[i] = 'a' + d
					} else {
						temp[i] = s[i]
					}
				} else {
					if i > 0 {
						temp[i] = s[i]
					} else {
						d := s[i] - 'a'
						temp[i] = 'A' + d
					}
				}
			}
			arr[j] = string(temp)
		}
	}
	var res string
	for i, v := range arr {
		if i != (len(arr) - 1) {
			res += v + " "
		} else {
			res += v
		}
	}

	return res
}
