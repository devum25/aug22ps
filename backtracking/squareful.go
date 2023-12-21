package backtracking

import "math"

// 996. Number of Squareful Arrays
// Hard
// 957
// 41
// Companies
// An array is squareful if the sum of every pair of adjacent elements is a perfect square.
// Given an integer array nums, return the number of permutations of nums that are squareful.
// Two permutations perm1 and perm2 are different if there is some index i such that perm1[i] != perm2[i].

// Example 1:

// Input: nums = [1,17,8]
// Output: 2
// Explanation: [1,8,17] and [17,8,1] are the valid permutations.
// Example 2:

// Input: nums = [2,2,2]
// Output: 1
func NumSquarefulPerms(nums []int) int {
	ans := make([][]int, 0)
	permutations(nums, 0, &ans)

	return len(ans)
}

func permutations(arr []int, pos int, ans *[][]int) {
	if len(arr) == pos {
		temp := make([]int, len(arr))

		for i := 0; i < len(temp); i++ {
			temp[i] = arr[i]
		}

		*ans = append(*ans, temp)
		return

	}
	hash := make(map[int]bool)
	for i := pos; i < len(arr); i++ {
		if _, ok := hash[arr[i]]; ok {
			continue
		}
		hash[arr[i]] = true
		arr[i], arr[pos] = arr[pos], arr[i]
		if pos == 0 || (pos > 0 && isPerfectSquare(arr[pos]+arr[pos-1])) {
			permutations(arr, pos+1, ans)
		}
		arr[i], arr[pos] = arr[pos], arr[i]

	}
}

func isPerfectSquare(x int) bool {
	// Find floating point value of
	// square root of x.
	if x >= 0 {

		val := math.Sqrt(float64(x))

		// if product of square root
		//is equal, then
		// return T/F
		if math.Mod(float64(val), float64(1)) != 0 {
			return false
		}
		return true
	}
	// else return false if n<0
	return false
}
