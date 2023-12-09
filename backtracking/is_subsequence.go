package backtracking

func IsSubsequence(s string, t string) bool {
	// lst := make([]byte, 0)
	// valid := false
	return solveOptimal(t, s)
	// Solve(t, s, 0, &lst, &valid)
	// return valid
}

func Solve(ip string, search string, idx int, lst *[]byte, valid *bool) {
	if *valid {
		return // If already found a match, no need to continue
	}
	x := *lst
	if search == string(x[:]) {
		*valid = true
	} else if idx == len(ip) {
		return
	}

	*lst = append(*lst, ip[idx])
	Solve(ip, search, idx+1, lst, valid)

	y := *lst
	y = y[:len(y)-1]
	*lst = y
	Solve(ip, search, idx+1, lst, valid)
}

func solveOptimal(input string, search string) bool {
	j := 0
	for i := 0; i < len(input); i++ {
		if input[i] == search[j] {
			j++
			if j == len(search) {
				break
			}
		}
	}

	if j == len(search) {
		return true
	}

	return false
}
