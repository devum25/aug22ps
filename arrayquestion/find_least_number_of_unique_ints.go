package arrayquestion

import "sort"

// Given an array of integers arr and an integer k. Find the least number of unique integers after removing exactly k elements.

type keyValue struct {
	key   int
	value int
}

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	hash := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		hash[arr[i]]++
	}

	// Extract key-value pairs from the map
	var pairs []keyValue
	for key, value := range hash {
		pairs = append(pairs, keyValue{key, value})
	}

	// Sort key-value pairs based on values in descending order
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	// Remove elements until k elements are removed
	for _, pair := range pairs {
		if k == 0 {
			break
		}
		if pair.value >= k {
			k = k - pair.value
			delete(hash, pair.key)
		}
	}

	count := len(hash)
	return count
}
