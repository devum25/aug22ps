package hashset

func LengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	ans := 0

	temp := 0
	startidx := 0
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; !ok {
			hash[s[i]] = i
			temp++
			if temp > ans {
				ans = temp
			}
		} else {
			if hash[s[i]] < startidx {
				hash[s[i]] = i
				temp++
				if temp > ans {
					ans = temp
				}
			} else if hash[s[i]] == (i - 1) {
				temp = 1
				startidx = i
				hash[s[i]] = i
			} else {
				temp = i - hash[s[i]]
				startidx = hash[s[i]] + 1
				hash[s[i]] = i
			}
		}
	}

	return ans
}
