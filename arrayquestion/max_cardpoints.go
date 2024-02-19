package arrayquestion

func MaxScore(cardPoints []int, k int) int {

	n := k
	ans := 0
	for i := 0; i < n; i++ {
		ans += cardPoints[i]
	}

	j := len(cardPoints) - 1

	i := k - 1
	sum := ans
	for i >= 0 && j >= (len(cardPoints)-k) {
		sum = sum - cardPoints[i]
		sum = sum + cardPoints[j]

		if sum > ans {
			ans = sum
		}
		i--
		j--
	}

	return ans
}
