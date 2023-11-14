package arrayquestion

// 485. Max Consecutive Ones
// Easy
// 5.3K
// 445
// Companies
// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Example 1:

// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
// Example 2:

// Input: nums = [1,0,1,1,0,1]
// Output: 2

func findMaxConsecutiveOnes(nums []int) int {

	count := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		} else if nums[i] == 0 {
			if temp > count {
				count = temp
			}
			temp = 0
		}
	}
	if temp > 0 && temp > count {
		count = temp
	}

	return count
}
