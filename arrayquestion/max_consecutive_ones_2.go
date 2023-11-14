package arrayquestion

// 485. Max Consecutive Ones - 3
// Easy
// 5.3K
// 445
// Companies
// Given a binary array nums, you are allowed to SWAP one of the 0 and 1
//return the maximum number of consecutive 1's in the array.

// Example 1:

// Input: nums = [1,1,0,1,1,0,1,1,1] --> [0,1,0,1,1,1,1,1,1]
// Output: 6

// Input: nums = [1,0,1,1,0,1]
// Output: 4

func MaxConsecutiveOneWithOneSwapAllowed(arr []int) int {
	ans := 0

	// find count of number of 1s and 0s
	one := 0
	zero := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			one++
		} else {
			zero++
		}
	}
	left := 0
	right := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j := i - 1
			k := i + 1
			for j >= 0 {
				if arr[j] == 1 {
					left++
				} else {
					break
				}
				j--
			}
			for k < len(arr) {
				if arr[k] == 1 {
					right++
				} else {
					break
				}
				k++
			}
			if left+right < one && zero > 0 {
				if ans < left+right+1 {
					ans = left + right + 1
				}
			} else {
				if left+right > ans {
					ans = left + right
				}
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
