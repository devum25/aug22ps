package recursion

import "math"

func KthGrammarOptimal(n, k int) int {
	if n == 1 && k == 1 {
		return 0
	}

	mid := int(math.Pow(2, float64(n-1))) / 2

	if k <= mid {
		return KthGrammarOptimal(n-1, k)
	} else {
		return 1 ^ KthGrammarOptimal(n-1, k-mid)
	}
}
