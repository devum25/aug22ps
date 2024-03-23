package greedy

import "sort"

func MaximumHappinessSum(happiness []int, k int) int64 {

	sort.Ints(happiness)

	i := 0
	j := len(happiness) - 1

	for i < j {
		happiness[i], happiness[j] = happiness[j], happiness[i]
		i++
		j--
	}

	i = 0
	var ans int64
	for i < len(happiness) && k > 0 {
		var x int64
		if (int64(happiness[i]) - int64(i)) > int64(0) {
			x = int64(happiness[i]) - int64(i)
		} else if (int64(happiness[i]) - int64(i)) <= 0 {
			x = 0
		} else {
			x = int64(happiness[i])
		}
		ans += x

		i++
		k--
	}

	return ans
}
