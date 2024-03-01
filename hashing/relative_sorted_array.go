package hashset

import "sort"

// Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
// Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.
// Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.

// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// Output: [2,2,2,1,4,3,3,9,6,7,19]
func RelativeSortArray(arr1 []int, arr2 []int) []int {

	hash := make(map[int]int)
	set := make(map[int]bool)

	for i := 0; i < len(arr2); i++ {
		set[arr2[i]] = true
	}
	res := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		if set[arr1[i]] {
			hash[arr1[i]]++
		} else {
			res = append(res, arr1[i])
		}
	}

	k := 0
	for i := 0; i < len(arr2); i++ {
		v := hash[arr2[i]]
		for v > 0 {
			arr1[k] = arr2[i]
			v--
			k++
		}
	}

	sort.Ints(res)
	i := 0
	for k < len(arr1) {
		arr1[k] = res[i]
		i++
		k++
	}

	return arr1
}

// Optinal Approach

func RelativeSortArrayOptimal(arr1 []int, arr2 []int) []int {
	index := 0
	for i := range arr2 {
		for j := range arr1 {
			if arr2[i] == arr1[j] {
				arr1[index], arr1[j] = arr1[j], arr1[index]
				index++
			}
		}
	}
	if index < len(arr1) {
		for i := index; i < len(arr1)-1; i++ {
			for j := i + 1; j < len(arr1); j++ {
				if arr1[i] > arr1[j] {
					arr1[i], arr1[j] = arr1[j], arr1[i]
				}
			}
		}
	}
	return arr1
}
