package arrayquestion

// A : 4, 3, 2, 7, 6, -2

func CountDirtyIndices(arr []int) int {
	ans := 0

	pfe := getPrefixSumOfEvenIndices(arr)
	pfo := getPrefixSumOfOddIndices(arr)

	for i := 0; i < len(arr); i++ {
		se := 0
		so := 0

		se = getSum(pfe, 0, i-1)
		se += getSum(pfo, i+1, len(arr)-1)

		so = getSum(pfo, 0, i-1)
		so += getSum(pfe, i+1, len(arr)-1)

		if so == se {
			ans++
		}
	}

	return ans
}

func getSum(arr []int, s, e int) int {
	if e < 0 {
		return 0
	}
	if s == 0 {
		return arr[e]
	} else if s > 0 && e < len(arr) {
		return arr[e] - arr[s-1]
	}
	return 0
}

func getPrefixSumOfEvenIndices(arr []int) []int {
	pfe := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			pfe[i] = arr[i]
			continue
		}
		if i%2 == 0 {
			pfe[i] = arr[i] + pfe[i-1]
		} else {
			pfe[i] = pfe[i-1]
		}
	}

	return pfe
}

func getPrefixSumOfOddIndices(arr []int) []int {
	pfo := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			pfo[i] = 0
			continue
		}
		if i%2 != 0 {
			pfo[i] = arr[i] + pfo[i-1]
		} else {
			pfo[i] = pfo[i-1]
		}
	}

	return pfo
}
