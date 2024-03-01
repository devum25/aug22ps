package arrayquestion

func Colorful(A int) []int {

	ans := make([]int, 0)
	res := make([]int, 0)
	for A > 0 {
		x := A % 10
		res = append(res, x)
		A = A / 10
	}

	i := 0
	j := len(res) - 1

	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	for i := 0; i < len(res); i++ {
		x := 1
		prev := res[i]
		for j := i; j < len(res); j++ {
			if i == j {
				ans = append(ans, res[i])
			} else {
				x = 10
				temp := prev*x + res[j]
				prev = temp
				ans = append(ans, temp)
			}
		}
	}

	return ans
}
