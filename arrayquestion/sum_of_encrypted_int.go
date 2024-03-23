package arrayquestion

func SumOfEncryptedInt(nums []int) int {

	for i := 0; i < len(nums); i++ {

		x := nums[i]
		t := 0
		c := 0
		for x > 0 {
			y := x % 10
			x = x / 10
			if y > t {
				t = y
			}
			c++
		}

		// x = nums[i]
		num := t
		c--
		p := 10
		for c > 0 {
			num = t + num*p
			c--
		}

		nums[i] = num

	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}

	return ans
}
