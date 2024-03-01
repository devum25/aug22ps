package hashset

// Given a number A, find if it is COLORFUL number or not.

// If number A is a COLORFUL number return 1 else, return 0.

// What is a COLORFUL Number:

// A number can be broken into different consecutive sequence of digits.
// The number 3245 can be broken into sequences like 3, 2, 4, 5, 32, 24, 45, 324, 245 and 3245.
// This number is a COLORFUL number, since the product of every consecutive sequence of digits is different

func Colorful(A int) bool {

	hash := make(map[int]bool)
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
		prev := res[i]
		for j := i; j < len(res); j++ {
			if i == j {
				if hash[res[i]] {
					return false
				}
				hash[res[i]] = true
			} else {
				temp := prev * res[j]
				if hash[temp] {
					return false
				}
				hash[temp] = true
				prev = temp
			}
		}
	}

	return true
}
