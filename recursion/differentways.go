package recursion

import "strconv"

func DiffWaysToCompute(expression string) []int {
	res := make([]int, 0)

	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			a := DiffWaysToCompute(expression[0:i])
			b := DiffWaysToCompute(expression[i+1:])

			for j := 0; j < len(a); j++ {
				for k := 0; k < len(b); k++ {
					if expression[i] == '+' {
						res = append(res, (a[j] + b[k]))
					} else if expression[i] == '-' {
						res = append(res, a[j]-b[k])
					} else if expression[i] == '*' {
						res = append(res, a[j]*b[k])
					}
				}
			}
		}
	}

	if len(res) == 0 {
		vl, _ := strconv.Atoi(expression)
		res = append(res, vl)
	}

	return res
}
