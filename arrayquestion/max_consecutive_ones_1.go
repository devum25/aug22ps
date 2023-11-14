package arrayquestion

// 485. Max Consecutive Ones - 1
// Easy
// 5.3K
// 445
// Companies
// Given a binary array nums, you are allowed to replace one of the 0 and 1
//return the maximum number of consecutive 1's in the array.

// Example 1:

// Input: nums = [1,1,0,1,1,0,1,1,1] --> [1,1,0,1,1,1,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
// Example 2:

// Input: nums = [1,0,1,1,0,1]
// Output: 2
func MaxConsecutiveOnesWithOneReplacementAllowed(arr []int) int {
	ans := 0
	left := 0
	right := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j := i - 1
			for j >= 0 {
				if arr[j] == 1 {
					left++
				} else {
					break
				}
				j--
			}
			k := i + 1
			for k < len(arr) {
				if arr[k] == 1 {
					right++
				} else {
					break
				}
				k++
			}

			if left+right+1 > ans {
				ans = left + right + 1
			}
			left = 0
			right = 0
		}
	}

	if ans == 0 {
		for i := 0; i < len(arr); i++ {
			ans++
		}
	}

	return ans
}
