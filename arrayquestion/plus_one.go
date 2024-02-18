package arrayquestion

func PlusOne(digits []int) []int {
	carry := 1

	res := make([]int, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		ans := digits[i] + carry
		carry = ans / 10
		if carry != 0 {
			res = append(res, ans%10)
		} else {
			res = append(res, ans)
		}
	}

	if carry != 0 {
		res = append(res, carry)
	}

	i := 0
	j := len(res) - 1

	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			break
		} else {
			res = res[i+1:]
		}
	}

	return res
}
