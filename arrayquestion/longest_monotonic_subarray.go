package arrayquestion

func longestMonotonicSubarray(nums []int) int {
	//[1,4,3,3,2]

	lar := make([]int, len(nums))
	small := make([]int, len(nums))

	lar[0] = 1
	small[0] = 1
	ans := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			small[i] = 0
			lar[i] = lar[i-1] + 1
			x := lar[i-1] + 1
			if x > ans {
				ans = x
			}
		} else if nums[i] > nums[i-1] {
			lar[i] = 0
			small[i] = small[i-1] + 1
			x := small[i-1] + 1
			if x > ans {
				ans = x
			}
		} else {
			lar[i] = 0
			small[i] = 0
		}

	}

	return ans
}
