package greedy

import "sort"

func BagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	score := 0

	i := 0
	j := len(tokens) - 1

	for i < j {
		if power >= tokens[i] {
			power = power - tokens[i]
			tokens[i] = 0
			score++
			i++
		} else {
			if score > 0 {
				score--
				power = power + tokens[j]
			}
			tokens[j] = 0
			j--
		}
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] > 0 {
			if power >= tokens[i] {
				power = power - tokens[i]
				score++
			}
		}
	}

	return score
}
