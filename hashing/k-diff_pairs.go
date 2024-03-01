package hashset

func FindPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	count := 0
	numFreq := make(map[int]int) // Store the frequency of each number

	for _, num := range nums {
		numFreq[num]++
	}

	for num := range numFreq {
		if k == 0 {
			// If k is 0, then we need to check if there are at least two occurrences of the number
			if numFreq[num] >= 2 {
				count++
			}
		} else {
			// Check if num + k or num - k exists in the map
			if _, ok := numFreq[num+k]; ok {
				count++
			}
			if _, ok := numFreq[num-k]; ok {
				count++
			}
		}
		// To ensure uniqueness, remove the number from the map
		delete(numFreq, num)
	}

	return count
}
